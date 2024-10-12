/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"pkg.amethysts.studio/renpy-graphviz/parser"
)

func main() {
	path, options := PlugCLI()

	content := parser.GetRenpyContent(path, options)

	graph, err := parser.Graph(content, options)
	if err != nil {
		parser.DocumentIssue(err)
	}

	err = graph.CreateFile("renpy-graphviz.dot")
	if err != nil {
		parser.DocumentIssue(err)
	}

	readDotFileToDrawGraph("renpy-graphviz.dot", "renpy-graphviz.png")

	if graph.Options.OpenFile {
		err = open.Run("renpy-graphviz.png")
		if err != nil {
			fmt.Println("A renpy-graphviz.png image file has been created, but couldn't be open. Please open it manually.")
		}
	}
}
