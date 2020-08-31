package muType2Spec

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/microtypes"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {

	fmt.Println("running muType2Spec")

	microList := &microtypes.MicroTypelist{
		MicroTypesByName:    map[string]*microtypes.MicroType{},
		MicroTypesASTByName: map[string]microtypes.MicroTypeAst{},
		MicroTypes:          []*microtypes.MicroType{},
	} // holds all muspecs
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(viper.GetStringSlice("importedTypeSpecs")...)

	globs := viper.GetStringSlice("muSpec.types")
	for _, glob := range globs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}

		LoadTypes(list, microList)

	}

	// build the new name and ast map
	for _, t := range microList.MicroTypes {
		regex := regexp.MustCompile(`^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Type)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", t.Type)
		}

		typeName := strings.TrimSpace(matches[1])
		microList.MicroTypesByName[typeName] = t
		microList.MicroTypesASTByName[typeName] = t.ToMicroTypeAst()
	}

	// update the typelist from microspecs
	microList.UpateTypelist(Typelist)

	Typelist.UpdateImports()
	typeAst.Format = viper.GetString("specFormat")
	Typelist.SaveAllTypeSpecsToDir(viper.GetString("typeSpecDir"))
}

func LoadTypes(list []string, bigList *microtypes.MicroTypelist) {

	for _, muSpecFile := range list {
		dataBytes, readError := ioutil.ReadFile(muSpecFile)
		if readError != nil {
			log.Fatal(readError)
		}
		microList := microtypes.MicroTypelist{}
		microList.Unmarshal(dataBytes)
		// add types to global list
		for _, mt := range microList.MicroTypes {
			bigList.MicroTypes = append(bigList.MicroTypes, mt)
		}
	}

}
