package main

import (
	"fmt"
	"hash/fnv"

	"github.com/emicklei/dot"
)

type Node struct {
	name      string
	neighbors []int
	repr      dot.Node
}

type RenpyGraph struct {
	nodes    map[int]*Node
	graphviz *dot.Graph
}

func NewGraph() RenpyGraph {
	g := dot.NewGraph(dot.Directed)
	return RenpyGraph{nodes: make(map[int]*Node), graphviz: g}
}

func (g RenpyGraph) PrettyPrint() {
	fmt.Println("\n====== Ren'Py Graph debug")
	fmt.Printf("%+v\n", g)

	for node := range g.nodes {
		fmt.Println(node, *g.nodes[node])
	}
	fmt.Println("=============")

}

func (g *RenpyGraph) AddNode(label string) {
	// fmt.Println("adding ", label, "to", g)
	_, ok := g.nodes[Hash(label)]
	if !ok {
		nodeGraph := g.graphviz.Node(label).Box()

		g.nodes[Hash(label)] = &Node{name: label, neighbors: make([]int, 0), repr: nodeGraph}
	}

}

func (g *RenpyGraph) AddEdge(label ...string) {

	g.graphviz.Edge(g.nodes[Hash(label[0])].repr, g.nodes[Hash(label[1])].repr).Label(label[2])

	g.nodes[Hash(label[0])].neighbors = append(g.nodes[Hash(label[0])].neighbors, Hash(label[1]))

}

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
