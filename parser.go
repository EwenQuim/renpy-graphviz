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
	Flowstop Situation = "Flowstop" // so here, Right = 3
)

func ParseRenPy(text []string) RenpyGraph {
	g := NewGraph()

	var lastLabel string

	context := Begin

	for _, line := range text {

		if matchesLabel, _ := regexp.MatchString(`renpy-graphviz.*BREAK`, line); matchesLabel {
			context = Situation(Flowstop)
		}
		// Label keyword
		matchesLabel, err := regexp.MatchString(`^\s*label .*:`, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesLabel {
			labelName := strings.TrimSpace(line)
			ignorelabelLabel, err := regexp.MatchString(`renpy-graphviz.*IGNORE`, line)
			if err != nil {
				fmt.Println(err)
			}
			if !ignorelabelLabel {
				labelName = labelName[6 : len(labelName)-1]

				g.AddNode(labelName)

				if lastLabel != "" && context == Situation(Label) {
					g.AddEdge(lastLabel, labelName, "label")
				}

				lastLabel = labelName
				context = Situation(Label)
			}

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
			context = Situation(Jump)
		}

	}

	return g
}
