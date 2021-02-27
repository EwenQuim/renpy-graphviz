package main

import (
	"fmt"
	"time"
)

func main() {
	defer track(runningtime("drawing"))

	path := PlugCLI()

	println(path)

	text := FileHandler(path[0])

	println(text)

	g := ParseRenPy(text)

	g.PrettyPrint()

	g.DrawNodes()

	//njbhibhub
	g.DrawEdges()

	for node := range g.node {
		fmt.Println(node, *g.node[node])
	}

	g.DrawGraph("test.png")

}

// Runningtime computes running time
func runningtime(s string) (string, time.Time) {
	fmt.Println("Start: ", s)
	return s, time.Now()
}

// Track is this
func track(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("End:   ", s, "took", endTime.Sub(startTime))
}
