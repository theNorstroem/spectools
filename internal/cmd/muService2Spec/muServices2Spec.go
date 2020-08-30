package muService2Spec

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/theNorstroem/spec-initializr/pkg/ast/serviceAst"
	"github.com/theNorstroem/spec-initializr/pkg/ast/typeAst"
	"github.com/theNorstroem/spec-initializr/pkg/microservices"

	"log"
	"path/filepath"

	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("running muService2Spec")
	microServicesList := &microservices.MicroServiceList{
		MicroServicesByName:    map[string]*microservices.MicroService{},
		MicroServicesASTByName: map[string]microservices.MicroServiceAst{},
		MicroServices:          []*microservices.MicroService{},
	} // holds all muspecs
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("serviceSpecDir"))
	Servicelist.LoadInstalledServiceSpecsFromDir(viper.GetStringSlice("importedServiceSpecs")...)

	globs := viper.GetStringSlice("muSpec.services")
	for _, glob := range globs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}

		LoadServices(list, microServicesList)

	}

	// build the new name and ast map
	for _, t := range microServicesList.MicroServices {

		serviceName := strings.TrimSpace(t.Package) + "." + strings.TrimSpace(t.Name)
		microServicesList.MicroServicesByName[serviceName] = t
		microServicesList.MicroServicesASTByName[serviceName] = t.ToMicroServiceAst()
	}

	// update the servicelist from microspecs
	//microServicesList.UpateServicelist(Servicelist)

	// types are needed for import checks
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(viper.GetStringSlice("importedTypeSpecs")...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))

	// update the services ast from microspecs
	microServicesList.UpateServicelist(Servicelist)

	Servicelist.UpdateAllImports(Typelist)

	typeAst.Format = viper.GetString("specFormat")
	Servicelist.SaveAllServiceSpecsToDir(viper.GetString("serviceSpecDir"))
}

func LoadServices(list []string, bigList *microservices.MicroServiceList) {

	for _, muSpecFile := range list {
		dataBytes, readError := ioutil.ReadFile(muSpecFile)
		if readError != nil {
			log.Fatal(readError)
		}
		microServicesList := microservices.MicroServiceList{}
		microServicesList.Unmarshal(dataBytes)
		// add types to global list
		for _, mt := range microServicesList.MicroServices {
			bigList.MicroServices = append(bigList.MicroServices, mt)
		}
	}

}
