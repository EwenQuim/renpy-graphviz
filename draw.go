package main

import (
	"bytes"
	"log"

	"github.com/goccy/go-graphviz"
)

func (g RenpyGraph) makeGraph() {

	if err := g.graphviz.RenderFilename(g.graph, graphviz.PNG, "renpy-graphviz.png"); err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := g.graphviz.Render(g.graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(buf.String())

}
