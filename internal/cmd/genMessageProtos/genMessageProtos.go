package genMessageProtos

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/protoTemplates"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/util"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

type singleTplData struct {
	Types     []specSpec.Type `json:"types,omitempty"`
	Imports   []string        `json:"imports,omitempty"`
	Package   string          `json:"package,omitempty"`
	preImport map[string]bool
	Options   map[string]string `json:"options,omitempty"`
}

func Run(cmd *cobra.Command, args []string) {
	// todo implement flag
	withInstalled := false

	fmt.Println("running genMessageProtos")
	allTypes := map[string]*specSpec.Type{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))

	for k, t := range Typelist.TypesByName {
		allTypes[k] = &t.TypeSpec
	}
	if withInstalled {
		for k, t := range Typelist.InstalledTypesByName {
			allTypes[k] = &t.TypeSpec
		}
	}

	protoTplData := map[string]*singleTplData{}

	// collect all types that belongs to one file
	for typeName, typeData := range allTypes {
		// typename is not the filename
		s := strings.Split(typeName, ".")
		filepath := strings.Join(s[:len(s)-1], "/") + "/" + typeData.XProto.Targetfile

		if protoTplData[filepath] == nil {
			protoTplData[filepath] = &singleTplData{
				preImport: map[string]bool{},
			}
		}

		protoTplData[filepath].Types = append(protoTplData[filepath].Types, *typeData)
		protoTplData[filepath].Package = typeData.XProto.Package
		protoTplData[filepath].Options = typeData.XProto.Options

		// pre imports
		for _, imp := range typeData.XProto.Imports {
			protoTplData[filepath].preImport[imp] = true
		}
		// sort types by name
		sort.Slice(protoTplData[filepath].Types, func(i, j int) bool {
			return protoTplData[filepath].Types[i].Name < protoTplData[filepath].Types[j].Name
		})
	}

	// process ipmports
	for _, tplData := range protoTplData {
		for imp, _ := range tplData.preImport {
			tplData.Imports = append(tplData.Imports, imp)
		}
	}

	// prepare templating
	fn := protoTemplates.GetSprigFuncs()
	tmpl, templateError := template.New("proto").Funcs(fn).Parse(protoTemplates.TypeTemplate)
	if templateError != nil {
		log.Fatal(templateError)
	}

	// clean the directory
	if viper.GetBool("proto.cleanBuild") {
		err := os.RemoveAll("./" + viper.GetString("build.proto.targetTypeDir")) // ./ is for safety purposes
		if err != nil {
			fmt.Println(err)
		}
	}
	// make the type protos
	for key, tplData := range protoTplData {
		filename := path.Join(viper.GetString("build.proto.targetTypeDir"), key)
		// create target dir => kann optimiert werden
		util.MkdirRelative(path.Dir(filename))

		sort.Strings(tplData.Imports)
		var buff bytes.Buffer
		// build proto from tpl
		err := tmpl.Execute(&buff, &tplData)
		if err != nil {
			log.Fatal(filename, err)
		}
		ioutil.WriteFile(filename, buff.Bytes(), 0644)
	}

}
