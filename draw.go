package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func MakeRenPyGraph(path string) {
	//

	text := fileHandler(path)

	g := parseRenPy(text)

	// Drawing the graph
	fmt.Println("Drawing the renpy-graphviz.png file...")

	if err := g.graphviz.RenderFilename(g.graph, graphviz.PNG, "renpy-graphviz.png"); err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := g.graphviz.Render(g.graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	writeFile("renpy-graphviz.dot", buf.String())

}

func GetDotGraph(content []string) string {
	g := parseRenPy(content)

	var buf bytes.Buffer
	if err := g.graphviz.Render(g.graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
