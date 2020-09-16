package genServiceProtos

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
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

type singleServiceTplData struct {
	Services  []specSpec.Service `json:"services,omitempty"`
	Imports   []string           `json:"imports,omitempty"`
	Package   string             `json:"package,omitempty"`
	preImport map[string]bool
	Options   map[string]string `json:"options,omitempty"`
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("genServiceProtos called")

	// todo implement flag
	withInstalled := false

	allServices := map[string]*specSpec.Service{}
	// types are needed for import checks
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))

	for k, t := range Servicelist.ServicesByName {
		allServices[k] = &t.ServiceSpec
	}
	if withInstalled {
		for k, t := range Servicelist.InstalledServicesByName {
			allServices[k] = &t.ServiceSpec
		}
	}

	// prepare templating
	fn := protoTemplates.GetSprigFuncs()
	tmpl, templateError := template.New("proto").Funcs(fn).Parse(protoTemplates.ServiceTemplate)
	if templateError != nil {
		log.Fatal(templateError)
	}

	// clean the directory
	if viper.GetBool("proto.cleanBuild") {
		err := os.RemoveAll("./" + viper.GetString("build.proto.targetDir"))
		if err != nil {
			fmt.Println(err)
		}
	}
	protoTplData := map[string]*singleServiceTplData{}

	// collect all types that belongs to one file
	for typeName, serviceData := range allServices {
		// typename is not the filename
		s := strings.Split(typeName, ".")
		filepath := strings.Join(s[:len(s)-1], "/") + "/" + serviceData.XProto.Targetfile

		if protoTplData[filepath] == nil {
			protoTplData[filepath] = &singleServiceTplData{
				preImport: map[string]bool{},
			}
		}

		protoTplData[filepath].Services = append(protoTplData[filepath].Services, *serviceData)
		protoTplData[filepath].Package = serviceData.XProto.Package
		protoTplData[filepath].Options = serviceData.XProto.Options

		// pre imports
		for _, imp := range serviceData.XProto.Imports {
			protoTplData[filepath].preImport[imp] = true
		}
		// sort types by name
		sort.Slice(protoTplData[filepath].Services, func(i, j int) bool {
			return protoTplData[filepath].Services[i].Name < protoTplData[filepath].Services[j].Name
		})
	}

	// process ipmports
	for _, tplData := range protoTplData {
		for imp, _ := range tplData.preImport {
			tplData.Imports = append(tplData.Imports, imp)
		}
	}

	// make the  protos
	for serviceName, tplData := range protoTplData {
		filename := path.Join(viper.GetString("build.proto.targetDir"), serviceName)
		// create target dir => kann optimiert werden
		util.MkdirRelative(path.Dir(filename))

		sort.Strings(tplData.Imports)

		var buff bytes.Buffer
		// build proto from tpl
		err := tmpl.Execute(&buff, &tplData)
		if err != nil {
			log.Fatal(serviceName, err)
		}
		ioutil.WriteFile(filename, buff.Bytes(), 0644)
	}

}
