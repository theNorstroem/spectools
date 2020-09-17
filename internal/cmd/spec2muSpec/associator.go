package spec2muSpec

import (
	"sync"
)

const (
	Service = iota + 1
	Type
	MuService
	MuType
)

type Graph struct {
	nodes    []*Node
	edges    map[*Node][]*Node // undirected
	edgesIn  map[*Node][]*Node
	edgesOut map[*Node][]*Node
	lock     sync.RWMutex
}

type Node struct {
	value interface{}
	Kind  int
}

func (g Graph) GetConnectedNodes(n *Node) []*Node {

	return g.edges[n]
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *Graph) AddUndirectedEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[*Node][]*Node)
	}
	if g.edgesOut == nil {
		g.edgesOut = make(map[*Node][]*Node)
	}
	if g.edgesIn == nil {
		g.edgesIn = make(map[*Node][]*Node)
	}
	g.edges[n1] = append(g.edges[n1], n2)
	g.edges[n2] = append(g.edges[n2], n1)
	g.lock.Unlock()
}

//  adds a directed edge to the graph. N1 -> N2
func (g *Graph) AddDirctedEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[*Node][]*Node)
	}
	if g.edgesOut == nil {
		g.edgesOut = make(map[*Node][]*Node)
	}
	if g.edgesIn == nil {
		g.edgesIn = make(map[*Node][]*Node)
	}
	g.edgesOut[n1] = append(g.edges[n1], n2)
	g.edgesIn[n2] = append(g.edges[n2], n1)
	g.lock.Unlock()
}
