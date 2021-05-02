package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// PlugCLI handles the Command Line Interface
func PlugCLI() (string, bool) {

	flag.Usage = func() {
		color.Set(color.Bold)
		fmt.Fprintf(os.Stderr, "Usage of: ")
		color.Blue(os.Args[0])
		color.Unset()
		fmt.Println("  args\n\tPath to your Ren'Py game folder")

		flag.PrintDefaults()

	}

	var labelsEdge bool

	flag.BoolVar(&labelsEdge, "e", false, "Do not display choice labels on edges")

	flag.Parse()

	if len(flag.Args()) == 0 {
		return ".", !labelsEdge
	}
	return flag.Args()[0], !labelsEdge

}
