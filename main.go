/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"github.com/skratchdot/open-golang/open"
	"pkg.amethysts.studio/renpy-graphviz/parser"
)

func main() {

	path, options := PlugCLI()

	content := parser.GetRenpyContent(path)

	graph, err := parser.Graph(content, options)
	if err != nil {
		parser.DocumentIssue(err)
	}

	graph.CreateFile("renpy-graphviz.dot")

	readDotFileToDrawGraph("renpy-graphviz.dot", "renpy-graphviz.png")

	if graph.Options.OpenFile {
		open.Run("renpy-graphviz.png")
	}

}
