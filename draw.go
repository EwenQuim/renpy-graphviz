package main

import (
	"fmt"

	"github.com/emicklei/dot"
)

func (g RenpyGraph) MakeGraph() {

	fmt.Println(g.graphviz.String())

	WriteFile("test.gv", g.graphviz.String())

}

func kkk() {
	g := dot.NewGraph(dot.Directed)
	n1 := g.Node("coding")
	n2 := g.Node("testing a little").Box()

	g.Edge(n1, n2)
	g.Edge(n2, n1, "back").Attr("color", "red")

	fmt.Println(g.String())
}
