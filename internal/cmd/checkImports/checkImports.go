package checkImports

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/util"
)

func Run(cmd *cobra.Command, args []string) {
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))

	Typelist.UpdateImports()

	typeAst.Format = viper.GetString("specFormat")
	Typelist.SaveAllTypeSpecsToDir(viper.GetString("specDir"))

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))

	Servicelist.UpdateAllImports(Typelist)

	serviceAst.Format = viper.GetString("specFormat")

	Servicelist.SaveAllServiceSpecsToDir(viper.GetString("specDir"))
}
