package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// PlugCLI handles the Command Line Interface
func PlugCLI() []string {

	flag.Usage = func() {
		color.Set(color.Bold)
		fmt.Fprintf(os.Stderr, "Usage of: ")
		color.Blue(os.Args[0])
		color.Unset()

		flag.PrintDefaults()

		fmt.Println("  args\n\tPath to your Ren'Py game folder")

	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		return []string{"."}
	} else {
		return flag.Args()
	}

}
