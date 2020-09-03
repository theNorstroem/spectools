package exportAsYaml

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/util"
	"gopkg.in/yaml.v3"
)

func Run(cmd *cobra.Command, args []string) {

	allTypes := map[string]*specSpec.Type{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {
		allTypes[k] = &t.TypeSpec
	}
	for k, t := range Typelist.InstalledTypesByName {
		allTypes[k] = &t.TypeSpec
	}

	allServices := map[string]*specSpec.Service{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("serviceSpecDir"))
	for k, t := range Servicelist.ServicesByName {
		allServices[k] = &t.ServiceSpec
	}
	for k, t := range Servicelist.InstalledServicesByName {
		allServices[k] = &t.ServiceSpec
	}

	output := map[string]interface{}{}

	output["types"] = allTypes
	output["services"] = allServices
	output["config"] = viper.AllSettings()

	outputstr, _ := yaml.Marshal(output)

	fmt.Print(string(outputstr))
}
