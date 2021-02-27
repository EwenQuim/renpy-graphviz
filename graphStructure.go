package main

type Node struct {
	name string
}

type RenpyGraph struct {
	edges map[Node][]Node
}

func NewGraph() RenpyGraph {
	return RenpyGraph{edges: make(map[Node][]Node)}
}
