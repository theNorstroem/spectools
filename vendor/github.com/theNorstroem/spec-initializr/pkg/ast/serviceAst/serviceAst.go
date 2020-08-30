package serviceAst

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spec-initializr/internal/specSpec"
	"github.com/theNorstroem/spec-initializr/pkg/ast/typeAst"
	"github.com/theNorstroem/spec-initializr/pkg/orderedmap"
	"github.com/theNorstroem/spec-initializr/pkg/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Servicelist struct {
	ServicesByName          map[string]*ServiceAst // this are the ones from the spec dir
	InstalledServicesByName map[string]*ServiceAst
	SpecDir                 string
}

type ServiceAst struct {
	Path        string // relative path of spec file to SpecDir
	FileName    string
	ServiceSpec specSpec.Service
}

var Format = "json"

// loads a spec directory and installed specs to the servicelist
func (l *Servicelist) LoadServiceSpecsFromDir(specDir string) {
	l.ServicesByName = loadServiceSpecsFromDir(specDir)
}

// loads a spec directory and installed specs to the servicelist
func (l *Servicelist) LoadInstalledServiceSpecsFromDir(specDir ...string) {
	// create map if it does not exist
	if l.InstalledServicesByName == nil {
		l.InstalledServicesByName = map[string]*ServiceAst{}
	}
	for _, dir := range specDir {
		tlist := loadServiceSpecsFromDir(dir)
		for tname, v := range tlist {
			l.InstalledServicesByName[tname] = v
		}
	}
}

func loadServiceSpecsFromDir(specDir string) (servicesMap map[string]*ServiceAst) {
	servicesMap = map[string]*ServiceAst{}
	err := filepath.Walk(specDir,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(fpath, "service.spec") {
				filename := path.Base(fpath)
				sdlen := len(strings.Split(specDir, "/"))
				if strings.HasPrefix(specDir, "./") {
					sdlen--
				}

				relativePath := path.Dir(strings.Join(strings.Split(fpath, "/")[sdlen:], "/"))
				AstService := &ServiceAst{
					Path:        relativePath, // store Path without specDir
					FileName:    filename,
					ServiceSpec: readAndUnmarshalSpec(fpath),
				}

				servicesMap[strings.Join([]string{AstService.ServiceSpec.XProto.Package, AstService.ServiceSpec.Name}, ".")] = AstService
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return servicesMap
}

func readAndUnmarshalSpec(fpath string) (s specSpec.Service) {
	dataBytes, readError := ioutil.ReadFile(fpath)
	if readError != nil {
		log.Fatal(readError)
	}
	parseError := yaml.Unmarshal(dataBytes, &s) //reads yaml and json because json is just a subset of yaml
	// convert fields from yaml.Node to  service
	if s.Services != nil {
		for pair := s.Services.Oldest(); pair != nil; pair = pair.Next() {
			yamlNode := pair.Value.(*yaml.Node)
			var RpcAST *specSpec.Rpc
			yamlNode.Decode(&RpcAST)
			pair.Value = RpcAST
			// query is also ordered
			if RpcAST.Query != nil {
				for pair := RpcAST.Query.Oldest(); pair != nil; pair = pair.Next() {
					queryYamlNode := pair.Value.(*yaml.Node)
					var query *specSpec.Queryparam
					queryYamlNode.Decode(&query)
					pair.Value = query
				}
			}
		}
	} else {
		// create empty object on nil
		s.Services = orderedmap.New()
	}

	if parseError != nil {
		log.Fatal(parseError)
	}
	return s
}

// check and updates the imports against a typlelistAst
func (l Servicelist) UpdateAllImports(typelist *typeAst.Typelist) {
	for _, s := range l.ServicesByName {
		s.UpdateImports(typelist)
	}
}

func (l *Servicelist) SaveAllServiceSpecsToDir(specDir string) {
	for _, serviceAst := range l.ServicesByName {
		serviceAst.Save(specDir)
	}
}

// save
func (ast ServiceAst) Save(specDir string) {
	filepath := path.Dir(path.Join(specDir, ast.Path, ast.FileName))
	filename := path.Join(filepath, ast.FileName)

	var d []byte
	var err error
	switch Format {
	case "json":
		d, err = ast.ToJson()
		break
	case "yaml":
		d, err = ast.ToYaml()
		break
	}
	if err != nil {
		panic(err)
	}
	util.MkdirRelative(filepath)
	err = ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *ServiceAst) ToJson() (d []byte, err error) {
	d, err = json.MarshalIndent(a.ServiceSpec, "", " ")
	return d, err
}

// returns unindented json
func (a *ServiceAst) ToJsonFlat() (d []byte, err error) {
	d, err = json.Marshal(a.ServiceSpec)
	return d, err
}

func (a *ServiceAst) ToYaml() (d []byte, err error) {
	d, err = yaml.Marshal(&a.ServiceSpec)
	return d, err
}

func (ast ServiceAst) UpdateImports(typelist *typeAst.Typelist) {
	// query.*.type
	ast.ServiceSpec.Services.Map(func(iKey interface{}, iValue interface{}) {
		rpc := iValue.(*specSpec.Rpc)
		rpc.Query.Map(func(qkey interface{}, qvalue interface{}) {
			qp := qvalue.(*specSpec.Queryparam)
			if qp.Type != "string" {
				imp, found := typelist.ResolveProtoImportForType(qp.Type)
				if found {
					// just add the imports, duplicates will be removed later
					ast.ServiceSpec.XProto.Imports = append(ast.ServiceSpec.XProto.Imports, imp)
				} else {

					fmt.Println(util.ScanForStringPosition(qp.Type,
						path.Join(viper.GetString("serviceSpecDir"), ast.FileName)), ":Import",
						qp.Type, "not found in Service",
						ast.ServiceSpec.Name, "on param", qkey.(string))
				}
			}
		})
		// data.request types
		// rpc.Data.Request
		if rpc.Data.Request != "" {
			imp, found := typelist.ResolveProtoImportForType(rpc.Data.Request)
			if found {
				// just add the imports, duplicates will be removed later
				ast.ServiceSpec.XProto.Imports = append(ast.ServiceSpec.XProto.Imports, imp)
			} else {

				fmt.Println(util.ScanForStringPosition(rpc.Data.Request,
					path.Join(viper.GetString("serviceSpecDir"), ast.FileName)), ":Import",
					rpc.Data.Request, "not found in Service",
					ast.ServiceSpec.Name, "on param", rpc.RpcName)
			}
		}

		// data.response types
		// rpc.Data.Response
		if rpc.Data.Response != "" {
			imp, found := typelist.ResolveProtoImportForType(rpc.Data.Response)
			if found {
				// just add the imports, duplicates will be removed later
				ast.ServiceSpec.XProto.Imports = append(ast.ServiceSpec.XProto.Imports, imp)
			} else {

				fmt.Println(util.ScanForStringPosition(rpc.Data.Response,
					path.Join(viper.GetString("serviceSpecDir"), ast.FileName)), ":Import",
					rpc.Data.Response, "not found in Service",
					ast.ServiceSpec.Name, "on param", rpc.RpcName)
			}
		}

		// remove duplicates
		keys := make(map[string]bool)
		list := []string{}
		for _, entry := range ast.ServiceSpec.XProto.Imports {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}

		ast.ServiceSpec.XProto.Imports = list

	})

	typelist.ResolveProtoImportForType("string")
}

// Deletes the spec from disk and removes the element from List
func (l *Servicelist) DeleteService(servicename string) {
	// delete the file
	filepath := path.Join(viper.GetString("serviceSpecDir"), l.ServicesByName[servicename].Path, l.ServicesByName[servicename].FileName)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DELETED", filepath)
	}

	delete(l.ServicesByName, servicename)

}
