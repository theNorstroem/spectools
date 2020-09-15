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

type ClientTypeList map[string]*clientspec.Type

func Run(cmd *cobra.Command, args []string) {

	allTypes := ClientTypeList{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("typeSpecDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {
		allTypes[k] = clientspec.CreateClientTypeFromAstType(&t.TypeSpec)
	}
	for k, t := range Typelist.InstalledTypesByName {
		allTypes[k] = clientspec.CreateClientTypeFromAstType(&t.TypeSpec)
	}
	CleanCPlusPlusStyleForTypes(allTypes)

	allServices := map[string]*clientspec.Service{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("serviceSpecDir"))
	for _, t := range Servicelist.ServicesByName {
		allServices[t.ServiceSpec.Name] = clientspec.CreateServiceFromAstService(&t.ServiceSpec)
	}
	for _, t := range Servicelist.InstalledServicesByName {
		allServices[t.ServiceSpec.Name] = clientspec.CreateServiceFromAstService(&t.ServiceSpec)
	}
	CleanCPlusPlusStyleForServices(allServices, Servicelist, allTypes)
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

func CleanCPlusPlusStyleForServices(services map[string]*clientspec.Service, servicelist *serviceAst.Servicelist, types ClientTypeList) {
	for _, service := range services {

		service.Services.Map(func(iKey interface{}, iValue interface{}) {
			servivename := iKey.(string)
			rpc := iValue.(*clientspec.CompressedService)
			//todo make a better resolution for services
			if strings.HasPrefix(rpc.Data.Request, ".") {
				// type starts from root, just remove the .
				rpc.Data.Request = rpc.Data.Request[1:len(rpc.Data.Request)]
			}
			if strings.HasPrefix(rpc.Data.Response, ".") {
				// type starts from root, just remove the .
				rpc.Data.Response = rpc.Data.Response[1:len(rpc.Data.Response)]
			}
			service.Services.Set(servivename, rpc)
		})

	}
}

func CleanCPlusPlusStyleForTypes(typeList ClientTypeList) {

	// https://developers.google.com/protocol-buffers/docs/proto3
	// Packages and Name Resolution
	// Type name resolution in the protocol buffer language works like C++: first the innermost scope is searched, then the next-innermost, and so on, with each package considered to be "inner" to its parent package. A leading '.' (for example, .foo.bar.Baz) means to start from the outermost scope instead.

	// The protocol buffer compiler resolves all type names by parsing the imported .proto files. The code generator for each language knows how to refer to each type in that language, even if it has different scoping rules.

	// resolves the c++ notation to client notation which is always from root

	// check all fields in all types
	for typename, typeordermap := range typeList {

		typeordermap.Fields.Map(func(iFname interface{}, iField interface{}) {
			fieldname := iFname.(string)
			iEnvField, _ := typeList[typename].Fields.Get(fieldname)
			es6Field := iEnvField.(*clientspec.Field)

			if strings.HasPrefix(es6Field.Type, ".") {
				// type starts from root, just remove the .
				es6Field.Type = es6Field.Type[1:len(es6Field.Type)]
			} else {
				// find first occurrence which can match the field

				es6Field.Type = cleanType(typename, es6Field.Type, typeList)

			}

			typeList[typename].Fields.Set(fieldname, es6Field)

		})

	}

}

func cleanType(typename string, fieldname string, typeList ClientTypeList) string {
	pathArr := strings.Split(typename, ".")
	// if we are in type a.b.c.d and want type x.y we look for
	// a.b.c.d.x.y
	// a.b.c.x.y
	// a.b.x.y
	// a.x.y
	// x.y
	for i := len(pathArr) - 1; i >= 0; i-- {
		sub := strings.Join(pathArr[0:i], ".")
		ftype := sub + "." + fieldname

		if typeList[ftype] != nil {
			// match
			return ftype
			i = 0
		}
		// we are at root
		if i == 0 && strings.HasPrefix(ftype, ".") {
			// remove .
			return ftype[1:len(ftype)]
		}
	}
	return fieldname
}
