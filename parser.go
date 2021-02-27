package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseRenPy(text []string) RenpyGraph {
	g := NewGraph()

	var lastLabel string

	for _, line := range text {
		fmt.Println(line)
		matchesLabel, err := regexp.MatchString(`^\s*label .*:`, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesLabel {
			labelName := strings.TrimSpace(line)
			labelName = labelName[6 : len(labelName)-1]
			lastLabel = labelName
			println("===== LABEL", labelName)

			g = g.AddNode(labelName)
		}

	}
	for _, line := range text {

		matchesJump, err := regexp.MatchString(`^\s*jump `, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesJump {
			jumpName := strings.TrimSpace(line)
			jumpName = jumpName[5:]

			println("===== JUMP", lastLabel, jumpName)
			g = g.AddEdge(lastLabel, jumpName)
		}

	}

	fmt.Println(g)
	return g
}
