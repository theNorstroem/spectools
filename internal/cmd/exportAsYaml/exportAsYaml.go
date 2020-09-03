package exportAsYaml

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/util"
	"gopkg.in/yaml.v3"
)

func Run(cmd *cobra.Command, args []string) {
	fullExport := false
	f := cmd.Flag("full")
	if f != nil {
		fullExport = f.Value.String() == "true"
	}

	allTypes := map[string]interface{}{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {
		if fullExport {
			allTypes[k] = t
		} else {
			allTypes[k] = t.TypeSpec
		}

	}
	for k, t := range Typelist.InstalledTypesByName {
		if fullExport {
			allTypes[k] = t
		} else {
			allTypes[k] = t.TypeSpec
		}
	}

	allServices := map[string]interface{}{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("serviceSpecDir"))
	for k, s := range Servicelist.ServicesByName {
		if fullExport {
			allServices[k] = s
		} else {
			allServices[k] = s.ServiceSpec
		}

	}
	for k, s := range Servicelist.InstalledServicesByName {
		if fullExport {
			allServices[k] = s
		} else {
			allServices[k] = s.ServiceSpec
		}
	}

	output := map[string]interface{}{}

	output["types"] = allTypes
	output["services"] = allServices
	output["config"] = viper.AllSettings()

	outputstr, _ := yaml.Marshal(output)

	fmt.Print(string(outputstr))
}
