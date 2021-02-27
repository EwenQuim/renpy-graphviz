package main

import (
	"fmt"
	"hash/fnv"
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type Node struct {
	name      string
	neighbors []int
	vizNode   *cgraph.Node
}

type RenpyGraph struct {
	node     map[int]*Node
	graphviz *graphviz.Graphviz
	drawing  *cgraph.Graph
}

func NewGraph() RenpyGraph {
	graphViz := graphviz.New()
	drawingGraph, err := graphViz.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := drawingGraph.Close(); err != nil {
			log.Fatal(err)
		}
		graphViz.Close()
	}()
	return RenpyGraph{node: make(map[int]*Node), graphviz: graphViz, drawing: drawingGraph}
}

func (g RenpyGraph) PrettyPrint() {
	fmt.Println("okkkkkkkk")
	fmt.Printf("%+v", g)
	fmt.Printf("%+v", g.drawing)
	fmt.Printf("%+v", g.graphviz)

	for node := range g.node {
		fmt.Println(node, *g.node[node])
	}
}

func (g *RenpyGraph) AddNode(label string) {
	// fmt.Println("adding ", label, "to", g)

	g.node[Hash(label)] = &Node{name: label, neighbors: make([]int, 0)}

}

func (g *RenpyGraph) AddEdge(labelFrom, labelTo string) {
	// fmt.Println("adding ", label, "to", g)

	g.node[Hash(labelFrom)].neighbors = append(g.node[Hash(labelFrom)].neighbors, Hash(labelTo))

}

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
