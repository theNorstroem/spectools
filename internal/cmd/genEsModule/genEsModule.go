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
	"strings"
)

func Run(cmd *cobra.Command, args []string) {

	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	clientspec.AddTypesToResolver(Typelist.TypesByName)
	clientspec.AddTypesToResolver(Typelist.InstalledTypesByName)
	// after adding all types we can build up the type resolutions
	clientspec.TransformCPlusStyleToAbsolutTypes()
	allTypes := clientspec.GetAllTypes()

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	clientspec.AddServicesToResolver(Servicelist.ServicesByName)
	clientspec.AddServicesToResolver(Servicelist.InstalledServicesByName)
	allServices := clientspec.GetAllServices()

	td, _ := json.Marshal(allTypes)

	// only escape " \n and \t and not just \ because we have strings like \u003c which is a <
	escapedType := strings.ReplaceAll(string(td), "\\\"", "\\\\\"")
	escapedType = strings.ReplaceAll(escapedType, "\\n", "\\\\n")
	escapedType = strings.ReplaceAll(escapedType, "\\t", "\\\\t")

	typeLine := "export const Types = JSON.parse(\n`" + escapedType + "`,\n);\n"
	sd, _ := json.Marshal(allServices)
	serviceLine := "export const Services = JSON.parse(\n`" + string(sd) + "`,\n);\n"

	err := ioutil.WriteFile(viper.GetString("build.esModule.targetFile"), []byte("/* eslint-disable */\n"+typeLine+"\n"+serviceLine), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Environment.js written to ", viper.GetString("build.esModule.targetFile"))

}
