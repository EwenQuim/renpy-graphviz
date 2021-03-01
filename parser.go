package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Situation string

const (
	Begin    Situation = "Begins"
	Label    Situation = "Label"    // if the first value in a constant block is `iota`
	Jump     Situation = "Jump"     // Go will automatically increment the rest for you.
	Filestop Situation = "Filestop" // so here, Right = 3
)

func ParseRenPy(text []string) RenpyGraph {
	g := NewGraph()

	var lastLabel string

	last := Begin

	for _, line := range text {

		if line == "FILESTOP" {
			last = Situation(Filestop)
		}
		// Label keyword
		matchesLabel, err := regexp.MatchString(`^\s*label .*:`, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesLabel {
			labelName := strings.TrimSpace(line)
			labelName = labelName[6 : len(labelName)-1]

			g.AddNode(labelName)

			if lastLabel != "" && last == Situation(Label) {
				g.AddEdge(lastLabel, labelName, "label")
			}

			lastLabel = labelName
			last = Situation(Label)
		}

		// Jump keyword
		matchesJump, err := regexp.MatchString(`^\s*jump `, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesJump {

			jumpName := strings.TrimSpace(line)
			jumpName = jumpName[5:]
			g.AddNode(jumpName)

			g.AddEdge(lastLabel, jumpName, "jump")
			last = Situation(Jump)
		}
		println(last)
	}

	return g
}
