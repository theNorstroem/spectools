package microservices

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/orderedmap"
	"github.com/theNorstroem/spectools/pkg/specSpec"
	"github.com/theNorstroem/spectools/pkg/util"
	"gopkg.in/yaml.v3"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type MicroServiceList struct {
	MicroServicesByName    map[string]*MicroService
	MicroServicesASTByName map[string]MicroServiceAst
	MicroServices          []*MicroService `yaml:"Services"`
}

// updates the core ast list
func (l *MicroServiceList) UpateServicelist(servicelist *serviceAst.Servicelist) {
	// build list to delete specs which are not Services.yaml
	deleteList := map[string]bool{}
	for serviceName, _ := range servicelist.ServicesByName {
		// mark every item as deletable
		deleteList[serviceName] = true
	}

	for serviceName, microServiceAst := range l.MicroServicesASTByName {
		deleteList[serviceName] = false
		// create type on Servicelist if it does not exist
		if servicelist.ServicesByName == nil {
			servicelist.ServicesByName = map[string]*serviceAst.ServiceAst{}
		}

		AstService, ok := servicelist.ServicesByName[serviceName]
		if !ok {
			servicelist.ServicesByName[serviceName] = &serviceAst.ServiceAst{
				Path:        microServiceAst.TargetPath,
				FileName:    microServiceAst.Name + ".service.spec",
				ServiceSpec: specSpec.Service{},
			}
			AstService = servicelist.ServicesByName[serviceName]
		}

		AstService.ServiceSpec.Name = microServiceAst.Name

		AstService.ServiceSpec.Description = microServiceAst.Description
		if AstService.ServiceSpec.XProto == nil {
			AstService.ServiceSpec.XProto = &specSpec.Typeproto{
				Imports:    []string{},
				Options:    map[string]string{},
				Package:    "",
				Targetfile: "",
			}
		}

		AstService.ServiceSpec.XProto.Package = microServiceAst.Package
		AstService.ServiceSpec.XProto.Targetfile = microServiceAst.Target
		AstService.ServiceSpec.XProto.Options["go_package"] = util.GetGoPackageName(microServiceAst.TargetPath)
		AstService.ServiceSpec.XProto.Options["java_package"] = "com." + microServiceAst.Package
		AstService.ServiceSpec.XProto.Options["java_outer_classname"] = strings.Title(strings.Replace(path.Base(microServiceAst.Target), ".proto", "Proto", 1))
		AstService.ServiceSpec.XProto.Options["java_multiple_files"] = "true"

		AstService.ServiceSpec.XProto.Imports = append(AstService.ServiceSpec.XProto.Imports, "google/api/annotations.proto")
		rpcServiceDeleteList := map[string]bool{}
		if AstService.ServiceSpec.Services != nil {
			AstService.ServiceSpec.Services.Map(func(iKey interface{}, iValue interface{}) {
				rpcServiceDeleteList[iKey.(string)] = true
			})
		}

		// Update the core rpc services with Ast services

		// update the delete list for rpc services
		microServiceAst.Services.Map(func(iKey interface{}, iValue interface{}) {
			rpcname := iKey.(string)
			rpcServiceDeleteList[rpcname] = false
		})
		// delete services in service (List, Create,...)
		for fieldname, del := range rpcServiceDeleteList {
			if del {
				AstService.ServiceSpec.Services.Delete(fieldname)
			}
		}

		// update each core service
		microServiceAst.Services.Map(func(iKey interface{}, iValue interface{}) {
			var sourceRPC *specSpec.Rpc
			var targetRPC *specSpec.Rpc
			rpcname := iKey.(string)
			sourceRPC = iValue.(*specSpec.Rpc)
			if AstService.ServiceSpec.Services == nil {
				AstService.ServiceSpec.Services = orderedmap.New()
			}
			m, found := AstService.ServiceSpec.Services.Get(rpcname)
			if found {
				// assign original when found
				targetRPC = m.(*specSpec.Rpc)

			} else {
				// create new if none was found in spec file
				targetRPC = &specSpec.Rpc{
					Data: &specSpec.Servicereqres{
						Request:  "google.protobuf.Empty",
						Response: "google.protobuf.Empty",
					},
					Deeplink: &specSpec.Servicedeeplink{
						Description: "Servicedeeplink desc",
						Href:        "",
						Method:      "GET",
						Rel:         "self",
					},
					Description: "",
					Query:       orderedmap.New(),
				}
			}

			// update target
			sourceRPC.Query.Map(func(iKey interface{}, iValue interface{}) {
				qpname := iKey.(string)
				s := iValue.(specSpec.Queryparam)
				targetRPC.Query.Set(qpname, &s)
			})

			targetRPC.RpcName = sourceRPC.RpcName
			targetRPC.Deeplink = sourceRPC.Deeplink
			targetRPC.Data.Request = sourceRPC.Data.Request
			targetRPC.Data.Response = sourceRPC.Data.Response
			targetRPC.Description = sourceRPC.Description
			AstService.ServiceSpec.Services.Set(rpcname, targetRPC)

		})

	}
	// delete the item
	for serviceName, del := range deleteList {
		if del {
			servicelist.DeleteService(serviceName)
		}
	}
}

type MicroRPC struct {
	Md string                 `yaml:"md"`
	Qp *orderedmap.OrderedMap `yaml:"qp"`
}

// holds a single service from microspec
type MicroService struct {
	Name        string     `yaml:"name"`
	Services    []MicroRPC `yaml:"services,omitempty"` //RPCs
	Target      string     `yaml:"target,omitempty"`
	Package     string     `yaml:"package,omitempty"`
	Description string     `yaml:"description"`
}

func (mt MicroService) ToMicroServiceAst() MicroServiceAst {

	// make the rpcMaps
	methods := orderedmap.New() //was map[string]RpcMap{}
	imports := []string{}

	imports = append(imports, "google/api/annotations.proto")

	// build the map

	for _, def := range mt.Services {
		// "List: GET /auth/{user} request.Type, response.Type #List eds with pagination"
		regex := regexp.MustCompile(`^([^:]+):\s?([A-Z]*)\s?([^\s]*) ?([^,]*)\s?,\s?([^#\s]*)\s?#?(.*)$`)
		matches := regex.FindStringSubmatch(def.Md)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", def.Md)
		}

		queryParams := orderedmap.New()

		def.Qp.Map(func(iKey interface{}, iValue interface{}) {

			var d string
			node := iValue.(*yaml.Node)
			node.Decode(&d)

			t := strings.Split(d, "#")

			qp := specSpec.Queryparam{
				Constraints: nil,
				Description: strings.TrimSpace(t[1]),
				Meta:        nil,
				Type:        strings.TrimSpace(t[0]),
			}
			queryParams.Set(iKey.(string), qp)
		})

		description := "No description available."
		if matches[6] != "" {
			description = matches[6]
		}

		r := &specSpec.Rpc{
			Data: &specSpec.Servicereqres{
				Request:  matches[4],
				Response: matches[5],
			},
			Deeplink: &specSpec.Servicedeeplink{
				Description: matches[0],
				Href:        matches[3],
				Method:      matches[2],
				Rel:         strings.ToLower(matches[1]),
			},
			Description: description,
			Query:       queryParams,
			RpcName:     matches[1] + mt.Name,
		}
		methods.Set(matches[1], r)

	}

	mAst := MicroServiceAst{
		Name:         mt.Name,
		Package:      mt.Package,
		Target:       mt.Target,
		Services:     methods,
		ProtoImports: imports, // not supported in service.yaml at the moment => empty list
		TargetPath:   strings.Join(strings.Split(mt.Package, "."), "/"),
		Description:  mt.Description,
	}

	return mAst
}

func NewRpcMap() RpcMap {
	return RpcMap{Readonly: false, Required: false, Repeated: false, Type: "string", DefaultValue: "", Description: "no description", FieldId: 1}
}

type MicroServiceAst struct {
	Name         string                 `yaml:"name"`
	ProtoImports []string               `yaml:"imports"`
	Package      string                 `yaml:"package,omitempty"`
	TargetPath   string                 // to find out the file to write
	Description  string                 `yaml:"description"`
	Services     *orderedmap.OrderedMap `yaml:"services,omitempty"` // with RPC
	Target       string                 `yaml:"target,omitempty"`
}

// field string will be converted to this type
// this type will be converted to rpcMap
type RpcMap struct {
	Readonly     bool
	Required     bool
	Repeated     bool
	Type         string
	DefaultValue string
	Description  string
	FieldId      int32
}

func (m *RpcMap) ParseServicestring(s string) {
	regex := regexp.MustCompile(`^(-*)?(\**)?(\[.?])?([^#=:]*):?([^=#]*)(=([^#]*))?(#(.*))?$`)
	matches := regex.FindStringSubmatch(s)
	if len(matches) == 0 {
		fmt.Println("field not parsed", s)
		return
	}
	if matches[1] != "" {
		m.Readonly = true
	}
	if matches[2] != "" {
		m.Required = true
	}
	if matches[3] != "" {
		m.Repeated = true
	}
	if matches[4] != "" {
		m.Type = strings.TrimSpace(matches[4])
	}
	if matches[5] != "" {
		n, _ := strconv.Atoi(strings.TrimSpace(matches[5]))
		m.FieldId = int32(n)
	} else {
		fmt.Println(util.ScanForStringPosition(s, viper.GetString("muSpec.Services"))+":Field numbers must be positive integers", s)
	}
	if matches[7] != "" {
		m.DefaultValue = strings.TrimSpace(matches[7])
	}
	if matches[9] != "" {
		m.Description = strings.TrimSpace(matches[9])
	}
}
