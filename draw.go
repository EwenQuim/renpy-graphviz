package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func DrawGraph(text string) {
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
	n, err := graph.CreateNode("n")
	if err != nil {
		log.Fatal(err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		log.Fatal(err)
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		log.Fatal(err)
	}
	e.SetLabel("e")

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
