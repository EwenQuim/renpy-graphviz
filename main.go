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

	MakeRenPyGraph(path[0])

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
