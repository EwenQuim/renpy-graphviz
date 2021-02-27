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
		// Label keyword
		matchesLabel, err := regexp.MatchString(`^\s*label .*:`, line)
		if err != nil {
			fmt.Println(err)
		}
		if matchesLabel {
			labelName := strings.TrimSpace(line)
			labelName = labelName[6 : len(labelName)-1]
			lastLabel = labelName

			g.AddNode(labelName)
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

			g.AddEdge(lastLabel, jumpName)
		}

	}

	return g
}
