package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"regexp"
	"strings"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// A Node is a REn'Py label and some properties, including its graph representation
type Node struct {
	name      string
	neighbors []int
	repr      *cgraph.Node
}

// RenpyGraph is the graph of Ren'Py story structure
type RenpyGraph struct {
	nodes    map[int]*Node
	graphviz *graphviz.Graphviz
	graph    *cgraph.Graph
}

// NewGraph creates a new graph
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

// AddNode to the ren'py graph, ignore if label already exists
func (g *RenpyGraph) AddNode(tags Tag, label string) {
	// fmt.Println("adding ", label, "to", g)
	re := regexp.MustCompile("[)_]")
	labelName := re.ReplaceAllString(label, " ")
	labelName = strings.Replace(labelName, "(", ": ", 1)

	_, ok := g.nodes[hash(label)]
	if !ok {
		nodeGraph, err := g.graph.CreateNode(label)
		if err != nil {
			log.Fatal(err)
		}
		nodeGraph.SetLabel(labelName)

		g.nodes[hash(label)] = &Node{name: label, neighbors: make([]int, 0), repr: nodeGraph}
	}
	if tags.title {
		g.nodes[hash(label)].repr.SetShape(cgraph.BoxShape).SetLabel(strings.ToUpper(labelName))
	}

}

// AddEdge to the renpy graph
func (g *RenpyGraph) AddEdge(tags Tag, label ...string) {

	parentNode := g.nodes[hash(label[0])]
	childrenNode := g.nodes[hash(label[1])]

	edge, err := g.graph.CreateEdge(parentNode.name+childrenNode.name, parentNode.repr, childrenNode.repr)
	if err != nil {
		log.Fatal(err)
	}

	if tags.lowLink {
		edge.SetStyle("dotted").SetColor("darkgreen")
	} else if tags.callLink {
		edge.SetStyle("dashed").SetColor("blue")
	}

	parentNode.neighbors = append(parentNode.neighbors, hash(label[1]))

}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
