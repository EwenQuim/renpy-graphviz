package parser

import (
	"fmt"
	"testing"
)

func (c *Context) String() string {
	str := ""
	if c.currentSituation != "" {
		str += fmt.Sprint(" situation:", c.currentSituation)
	}
	if c.currentLabel != "" {
		str += fmt.Sprint(" label:", c.currentLabel)
	}
	if c.lastLabel != "" {
		str += fmt.Sprint(" lastLabel:", c.lastLabel)
	}
	if c.linkedToLastLabel {
		str += " linked to last label"
	}

	return str
}

func (g RenpyGraph) testGraphEquality(f RenpyGraph, t *testing.T) {
	for nodeName, node := range g.nodes {
		fNode, ok := f.nodes[nodeName]
		if !ok {
			t.Errorf("Node '%v' wasn't expected to be generated", nodeName)
		}
		if node.name != fNode.name {
			t.Errorf("Node names '%v' and '%v' doesn't match", node.name, fNode.name)
		}
		for i, n := range node.neighbors {
			if n != fNode.neighbors[i] {
				t.Errorf("%v and %v don't match", node.neighbors, fNode.neighbors)
			}
		}
	}
	for nodeName := range f.nodes {
		_, ok := f.nodes[nodeName]
		if !ok {
			t.Errorf("Node '%v' was expected to be generated but wasn't", nodeName)
		}
	}
}
