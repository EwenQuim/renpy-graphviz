package parser

import (
	"testing"
)

// real tests for utils functions

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

func TestCleanContextAccordingToIndent(t *testing.T) {
	// Removes last
	c := Context{}
	c.indent = 5
	c.labelStack = []labelStack{
		{0, "first"},
		{2, "second"},
		{6, "third"},
	}
	c.cleanContextAccordingToIndent("nothing")

	if len(c.labelStack) != 2 {
		t.Errorf("labelStack length: expected %v, got %v", 2, c.labelStack)
	}
	if c.labelStack[len(c.labelStack)-1].labelName != "second" {
		t.Error("labelStack element unexpected", c.labelStack)
	}
	if c.lastLabel != "third" {
		t.Error("lastLabel unexpected", c.labelStack)
	}

	// Removes 2 last
	c.indent = 2
	c.labelStack = []labelStack{
		{0, "first"},
		{2, "second"},
		{6, "third"},
	}
	c.cleanContextAccordingToIndent("nothing")

	if len(c.labelStack) != 1 {
		t.Errorf("labelStack length: expected %v, got %v", 1, c.labelStack)
	}
	if c.labelStack[len(c.labelStack)-1].labelName != "first" {
		t.Error("labelStack element unexpected", c.labelStack)
	}
	if c.lastLabel != "second" {
		t.Error("lastLabel unexpected", c.labelStack)
	}

	// Removes everything
	c.indent = 0
	c.labelStack = []labelStack{
		{0, "first"},
		{2, "second"},
		{6, "third"},
	}
	c.cleanContextAccordingToIndent("nothing")
	if len(c.labelStack) != 0 {
		t.Errorf("labelStack length: expected %v, got %v", 0, c.labelStack)
	}
	if c.lastLabel != "first" {
		t.Error("lastLabel unexpected", c.labelStack)
	}

	// Screen state cleaning
	c.indent = 0
	c.currentScreen = "this_is_a_screen"
	c.cleanContextAccordingToIndent("nothing")
	if c.currentScreen != "" {
		t.Errorf("labelStack screen not cleaned")
	}
}
