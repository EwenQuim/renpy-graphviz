package main

import (
	"hash/fnv"

	"github.com/goccy/go-graphviz/cgraph"
)

type Node struct {
	name      string
	vizNode   *cgraph.Node
	neighbors []int
}

type RenpyGraph struct {
	node map[int]*Node
}

func NewGraph() RenpyGraph {
	return RenpyGraph{node: make(map[int]*Node)}
}

func (g RenpyGraph) AddNode(label string) RenpyGraph {
	// fmt.Println("adding ", label, "to", g)

	g.node[Hash(label)] = &Node{label, nil, make([]int, 0)}
	return g
}

func (g RenpyGraph) AddEdge(labelFrom, labelTo string) RenpyGraph {
	// fmt.Println("adding ", label, "to", g)

	g.node[Hash(labelFrom)].neighbors = append(g.node[Hash(labelFrom)].neighbors, Hash(labelTo))

	return g
}

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
