package main

import (
	"fmt"
	"hash/fnv"
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// A Node is a REn'Py label and some properties, including its graph representation
type Node struct {
	name      string
	neighbors []int
	repr      *cgraph.Node
}

// The graph of Ren'Py story structure
type RenpyGraph struct {
	nodes    map[int]*Node
	graphviz *graphviz.Graphviz
	graph    *cgraph.Graph
}

//
func NewGraph() RenpyGraph {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	return RenpyGraph{nodes: make(map[int]*Node), graphviz: g, graph: graph}
}

// PrettyPrint prints the graph in the terminal
func (g RenpyGraph) PrettyPrint() {
	fmt.Println("\n====== Ren'Py Graph debug")
	fmt.Printf("%+v\n", g)

	for node := range g.nodes {
		fmt.Println(node, *g.nodes[node])
	}
	fmt.Println("=============")

}

// AddNode to the ren'py graph
func (g *RenpyGraph) AddNode(label string) {
	// fmt.Println("adding ", label, "to", g)
	_, ok := g.nodes[hash(label)]
	if !ok {
		nodeGraph, err := g.graph.CreateNode(label)
		if err != nil {
			log.Fatal(err)
		}

		g.nodes[hash(label)] = &Node{name: label, neighbors: make([]int, 0), repr: nodeGraph}
	}

}

// AddEdge to the repy graph
func (g *RenpyGraph) AddEdge(label ...string) {

	edge, err := g.graph.CreateEdge(g.nodes[hash(label[0])].name+g.nodes[hash(label[1])].name, g.nodes[hash(label[0])].repr, g.nodes[hash(label[1])].repr)
	if err != nil {
		log.Fatal(err)
	}
	edge.SetLabel(label[2])

	g.nodes[hash(label[0])].neighbors = append(g.nodes[hash(label[0])].neighbors, hash(label[1]))

}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
