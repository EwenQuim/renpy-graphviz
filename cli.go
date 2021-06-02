package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	toml "github.com/pelletier/go-toml"
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
	var silent bool
	var openFile bool
	var hideScreens bool
	var hideNestedScreens bool
	var fullDebug bool

	// TOML
	getDefaultConfig()
	config, err := toml.LoadFile("renpy-graphviz.config")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		showAtoms = config.Get("config.atoms").(bool)
		hideEdgesLabels = !config.Get("config.edges").(bool)
		openFile = config.Get("config.open").(bool)
		hideScreens = !config.Get("config.screens").(bool)
		hideNestedScreens = !config.Get("config.nested-screens").(bool)
		silent = config.Get("config.silent").(bool)
		fullDebug = config.Get("config.debug").(bool)
	}

	// CLI overrides TOML
	flag.BoolVar(&showAtoms, "atoms", showAtoms, "Show atoms (lonely nodes)")
	flag.BoolVar(&hideEdgesLabels, "hide-edges", hideEdgesLabels, "Hide choice labels on edges")
	flag.BoolVar(&silent, "silent", silent, "Display nothing to the stdout")
	flag.BoolVar(&openFile, "open", openFile, "Open file in default image viewer")
	flag.BoolVar(&hideScreens, "hide-screens", hideScreens, "Hide screens")
	flag.BoolVar(&hideNestedScreens, "hide-nested", hideNestedScreens, "Hide nested screens")
	flag.BoolVar(&fullDebug, "debug", fullDebug, "Debug")

	flag.Parse()

	path := "."
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	}
	return path, parser.RenpyGraphOptions{
		ShowEdgesLabels:   !hideEdgesLabels,
		ShowAtoms:         showAtoms,
		Silent:            silent,
		OpenFile:          openFile,
		ShowScreens:       !hideScreens,
		ShowNestedScreens: !hideNestedScreens,
		FullDebug:         fullDebug}
}
