/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//defer track(runningtime("drawing"))

	path := PlugCLI()

	text := FileHandler(path[0])

	g := ParseRenPy(text)

	// g.PrettyPrint()

	g.MakeGraph()

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
