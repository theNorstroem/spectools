package deprecated

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/util"
)

func Run(cmd *cobra.Command, args []string) {
	installedTypes := map[string]interface{}{}
	allTypes := map[string]specSpec.Type{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {

		allTypes[k] = t.TypeSpec

	}

	for k, t := range Typelist.InstalledTypesByName {
		installedTypes[k] = t.TypeSpec
		allTypes[k] = t.TypeSpec
	}

	allServices := map[string]specSpec.Service{}
	installedServices := map[string]interface{}{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	for k, s := range Servicelist.ServicesByName {

		allServices[k] = s.ServiceSpec

	}

	for k, s := range Servicelist.InstalledServicesByName {
		installedServices[k] = s.ServiceSpec
		allServices[k] = s.ServiceSpec

	}

	// check every type in typelist and servicelist

	for tname, ast := range Typelist.TypesByName {
		ast.TypeSpec.Fields.Map(func(iKey interface{}, iValue interface{}) {
			f := iValue.(*specSpec.Field) //*string:1 # A * before the type means required
			if allTypes[f.Type].Lifecycle != nil && allTypes[f.Type].Lifecycle.Deprecated {
				fmt.Println("WARNING: field", iKey.(string), "in type", tname, "uses deprecated type", f.Type)
				fmt.Println(allTypes[f.Type].Lifecycle.Info)
			}
		})
	}
	for sname, _ := range Servicelist.ServicesByName {

		if 1 == 2 {
			fmt.Println(sname)
		}

		//fmt.Println(ast)
	}
}
