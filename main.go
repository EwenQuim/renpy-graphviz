/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"fmt"

	"pkg.amethysts.studio/renpy-graphviz/parser"
)

func main() {

	path, labelEdges := PlugCLI()

	content := parser.GetRenpyContent(path)

	graph := parser.Graph(content, labelEdges)

	graph.CreateFile("renpy-graphviz.dot")

	readDotFileToDrawGraph("renpy-graphviz.dot", "renpy-graphviz.png")

	fmt.Println("Done.")

}
