package genEsModule

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/clientspec"
	"github.com/theNorstroem/spectools/pkg/util"
	"io/ioutil"
)

func Run(cmd *cobra.Command, args []string) {

	allTypes := map[string]*clientspec.Type{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {
		allTypes[k] = clientspec.CreateFromAstType(&t.TypeSpec)
	}
	for k, t := range Typelist.InstalledTypesByName {
		allTypes[k] = clientspec.CreateFromAstType(&t.TypeSpec)
	}

	allServices := map[string]*clientspec.Service{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("serviceSpecDir"))
	for _, t := range Servicelist.ServicesByName {
		allServices[t.ServiceSpec.Name] = clientspec.CreateSectviceFromAstService(&t.ServiceSpec)
	}
	for _, t := range Servicelist.InstalledServicesByName {
		allServices[t.ServiceSpec.Name] = clientspec.CreateSectviceFromAstService(&t.ServiceSpec)
	}

	td, _ := json.Marshal(allTypes)
	typeLine := "export const Types = JSON.parse(`" + string(td) + "`)"
	sd, _ := json.Marshal(allServices)
	serviceLine := "export const Services = JSON.parse(`" + string(sd) + "`)"

	err := ioutil.WriteFile(viper.GetString("build.esModule.targetFile"), []byte(typeLine+"\n"+serviceLine), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Environment.js written to ", viper.GetString("build.esModule.targetFile"))

}
