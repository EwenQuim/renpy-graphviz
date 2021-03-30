package main

import (
	"io/ioutil"
	"log"

	"github.com/goccy/go-graphviz"
)

func ReadDotFileToDrawGraph(pathToDotfile, imageName string) {
	g := graphviz.New()
	b, err := ioutil.ReadFile(pathToDotfile)
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
