package spec2muSpec

// Generate µ-specs from Specs
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/internal/cmd/muSpec2Spec"
	"github.com/theNorstroem/spectools/pkg/ast/serviceAst"
	"github.com/theNorstroem/spectools/pkg/ast/typeAst"
	"github.com/theNorstroem/spectools/pkg/microservices"
	"github.com/theNorstroem/spectools/pkg/microtypes"
	"github.com/theNorstroem/spectools/pkg/util"
	"log"
	"path/filepath"
)

func Run(cmd *cobra.Command, args []string) {
	deleteMuSpecs := false
	f := cmd.Flag("delete")
	if f != nil {
		deleteMuSpecs = f.Value.String() == "true"
	}

	fmt.Println("running spec2muSpec")

	microTypesList := &microtypes.MicroTypelist{
		MicroTypesByName:    map[string]*microtypes.MicroType{},
		MicroTypesASTByName: map[string]microtypes.MicroTypeAst{},
		MicroTypes:          []*microtypes.MicroType{},
	} // holds all muspecs

	microServicesList := &microservices.MicroServiceList{
		MicroServicesByName:    map[string]*microservices.MicroService{},
		MicroServicesASTByName: map[string]microservices.MicroServiceAst{},
		MicroServices:          []*microservices.MicroService{},
	} // holds all muspecs

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	Servicelist.LoadInstalledServiceSpecsFromDir(viper.GetStringSlice("importedServiceSpecs")...)

	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	serviceglobs := viper.GetStringSlice("muSpec.services")
	typeglobs := viper.GetStringSlice("muSpec.types")
	for _, glob := range typeglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		muSpec2Spec.LoadTypes(list, microTypesList)
	}
	for _, glob := range serviceglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		muSpec2Spec.LoadServices(list, microServicesList)
	}

	var graph Graph
	a := &Node{
		value: "type a",
		Kind:  Service,
	}
	b := &Node{
		value: "mutype a",
		Kind:  Service,
	}
	c := &Node{
		value: "Entity a",
		Kind:  Type,
	}
	d := &Node{
		value: "Collection a",
		Kind:  Type,
	}
	graph.AddNode(a)
	graph.AddNode(b)
	graph.AddNode(c)
	graph.AddDirctedEdge(a, c)
	graph.AddDirctedEdge(a, d)
	graph.AddUndirectedEdge(b, a)

	l := graph.GetConnectedNodes(a)
	fmt.Println(l)

	fmt.Println(deleteMuSpecs)
}
