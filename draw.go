package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func (g RenpyGraph) DrawGraph(filename string) {

	var buf bytes.Buffer
	if err := g.graphviz.Render(g.drawing, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 3. write to file directly
	if err := g.graphviz.RenderFilename(g.drawing, graphviz.PNG, filename); err != nil {
		log.Fatal(err)
	}
}

func (g *RenpyGraph) DrawNodes() {

	for _, node := range g.node {
		g.drawing.CreateNode(node.name)

	}
}

func (g *RenpyGraph) DrawEdges() {

	for _, nodeParent := range g.node {
		for _, nodeChild := range nodeParent.neighbors {
			e, err := g.drawing.CreateEdge("jump", nodeParent.vizNode, g.node[nodeChild].vizNode)
			if err != nil {
				log.Fatal(err)
			}
			e.SetLabel("jump")

			print(e)

		}
	}
}
