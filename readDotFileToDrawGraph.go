package main

import (
	"io/ioutil"
	"log"

	"github.com/ewenquim/renpy-graphviz/parser"
	"github.com/goccy/go-graphviz"
)

func readDotFileToDrawGraph(pathToDotfile, imageName string) {
	defer parser.Track(parser.RunningTime("Drawing .png file"))

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
