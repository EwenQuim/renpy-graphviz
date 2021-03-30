/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"fmt"
	"time"

	"github.com/ewenquim/renpy-graphviz/parser"
)

func main() {
	defer track(runningtime("drawing"))

	path := PlugCLI()

	content := getRenpyContent(path)

	graph := parser.Graph(content)

	graph.CreateFile("renpy-graphviz.dot")

	ReadDotFileToDrawGraph("renpy-graphviz.dot", "renpy-graphviz.png")

	fmt.Println("Done.")

}

// Runningtime computes running time
func runningtime(s string) (string, time.Time) {
	return s, time.Now()
}

// Track is this
func track(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println(s, "took", endTime.Sub(startTime))
}
