package main

import (
	"fmt"
)

func (g RenpyGraph) MakeGraph() {

	fmt.Println(g.graphviz.String())

	WriteFile("test.gv", g.graphviz.String())

}
