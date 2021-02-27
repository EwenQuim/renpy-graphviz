package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func DrawGraph(renpyGraph RenpyGraph) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	// Draw nodes
	for index, node := range renpyGraph.node {
		renpyGraph.node[index].vizNode, err = graph.CreateNode(node.name)
		if err != nil {
			log.Fatal(err)
		}
		println(node.vizNode)

	}

	fmt.Println(" YOOOO", renpyGraph)
	println("edges")

	for _, nodeParent := range renpyGraph.node {
		println(nodeParent.vizNode)
		for _, nodeChild := range nodeParent.neighbors {
			e, err := graph.CreateEdge("jump", nodeParent.vizNode, renpyGraph.node[nodeChild].vizNode)
			if err != nil {
				log.Fatal(err)
			}
			e.SetLabel("jump")

		}
	}

	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "graph.png"); err != nil {
		log.Fatal(err)
	}
}
