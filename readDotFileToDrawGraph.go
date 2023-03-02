package main

import (
	"log"
	"os"

	"github.com/goccy/go-graphviz"
)

func readDotFileToDrawGraph(pathToDotfile, imageName string) {
	g := graphviz.New()
	b, err := os.ReadFile(pathToDotfile)
	if err != nil {
		log.Fatal(err)
	}
	graph, err := graphviz.ParseBytes(b)
	if err != nil {
		log.Fatal(err)
	}

	if err := g.RenderFilename(graph, graphviz.PNG, imageName); err != nil {
		log.Fatal(err)
	}
}
