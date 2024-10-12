package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	toml "github.com/pelletier/go-toml"
	"pkg.amethysts.studio/renpy-graphviz/parser"
)

// isFlag checks if a string is a flag (starts with a dash)
func isFlag(arg string) bool {
	return len(arg) > 0 && arg == "-"
}

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
	var skipFilesRegex string

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
		skipFilesRegex = config.Get("config.skip-files").(string)
	}

	// CLI overrides TOML
	flag.BoolVar(&showAtoms, "atoms", showAtoms, "Show atoms (lonely nodes)")
	flag.BoolVar(&hideEdgesLabels, "hide-edges", hideEdgesLabels, "Hide choice labels on edges")
	flag.BoolVar(&silent, "silent", silent, "Display nothing to the stdout")
	flag.BoolVar(&openFile, "open", openFile, "Open file in default image viewer")
	flag.BoolVar(&hideScreens, "hide-screens", hideScreens, "Hide screens")
	flag.BoolVar(&hideNestedScreens, "hide-nested", hideNestedScreens, "Hide nested screens")
	flag.BoolVar(&fullDebug, "debug", fullDebug, "Debug")
	flag.StringVar(&skipFilesRegex, "skip-files", skipFilesRegex, "Regex pattern for excluding files")

	// Manually handle the non-flag argument
	var path string
	if len(os.Args) > 1 && !isFlag(os.Args[1]) {
		path = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...) // Remove the non-flag argument from os.Args
	} else {
		path = "."
	}
	flag.Parse()
	return path, parser.RenpyGraphOptions{
		ShowEdgesLabels:   !hideEdgesLabels,
		ShowAtoms:         showAtoms,
		Silent:            silent,
		OpenFile:          openFile,
		ShowScreens:       !hideScreens,
		ShowNestedScreens: !hideNestedScreens,
		FullDebug:         fullDebug,
		SkipFilesRegex:    skipFilesRegex,
	}
}
