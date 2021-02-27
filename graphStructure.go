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

func (g *RenpyGraph) AddEdge(labelFrom, labelTo string) {
	// fmt.Println(g.nodes[Hash(labelFrom)])
	// fmt.Println(g.nodes[Hash(labelTo)])

	g.graphviz.Edge(g.nodes[Hash(labelFrom)].repr, g.nodes[Hash(labelTo)].repr)

	g.nodes[Hash(labelFrom)].neighbors = append(g.nodes[Hash(labelFrom)].neighbors, Hash(labelTo))

}

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
