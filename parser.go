package main

import (
	"fmt"
	"regexp"
	"strings"
)

type situation string

const (
	situationBegin    situation = "Begins"
	situationLabel    situation = "Label"
	situationJump     situation = "Jump"
	situationFlowstop situation = "Flowstop"
)

func parseRenPy(text []string) RenpyGraph {
	g := NewGraph()

	var lastLabel string

	context := situationBegin

	for _, line := range text {

		if matchesLabel, _ := regexp.MatchString(`renpy-graphviz.*BREAK`, line); matchesLabel {
			context = situationFlowstop
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

				if lastLabel != "" && context == situationLabel {
					g.AddEdge(lastLabel, labelName, "label")
				}

				lastLabel = labelName
				context = situationLabel
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

			g.AddEdge(lastLabel, jumpName, "")
			context = situationJump
		}

	}

	return g
}
