package exportAsYaml

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/clientspec"
	"github.com/theNorstroem/spectools/pkg/util"
	"gopkg.in/yaml.v3"
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
	for k, t := range Servicelist.ServicesByName {
		allServices[k] = clientspec.CreateSectviceFromAstService(&t.ServiceSpec)
	}
	for k, t := range Servicelist.InstalledServicesByName {
		allServices[k] = clientspec.CreateSectviceFromAstService(&t.ServiceSpec)
	}

	output := map[string]interface{}{}

	output["types"] = allTypes
	output["services"] = allServices
	outputstr, _ := yaml.Marshal(output)

	fmt.Print(string(outputstr))
}
