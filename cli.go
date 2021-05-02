package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"pkg.amethysts.studio/renpy-graphviz/parser"
)

// PlugCLI handles the Command Line Interface
func PlugCLI() (string, parser.RenpyGraphOptions) {

	flag.Usage = func() {
		color.Set(color.Bold)
		fmt.Fprintf(os.Stderr, "Usage of: ")
		color.Blue(os.Args[0])
		color.Unset()
		fmt.Println("  args\n\tPath to your Ren'Py game folder")

		flag.PrintDefaults()

	}

	var hideEdgesLabels bool
	var showAtoms bool

	flag.BoolVar(&showAtoms, "a", false, "Show atoms (lonely nodes)")
	flag.BoolVar(&hideEdgesLabels, "e", false, "Hide choice labels on edges")

	flag.Parse()

	if len(flag.Args()) == 0 {
		return ".", parser.RenpyGraphOptions{ShowEdgesLabels: !hideEdgesLabels, ShowAtoms: showAtoms}
	}
	return flag.Args()[0], parser.RenpyGraphOptions{ShowEdgesLabels: !hideEdgesLabels, ShowAtoms: showAtoms}

}
