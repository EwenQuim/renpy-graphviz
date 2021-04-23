package parser

import "testing"

func TestEmptyGraph(t *testing.T) {
	t.Parallel()

	graph := Graph([]string{"nothing"})

	expectedGraph := RenpyGraph{}

	graph.testGraphEquality(expectedGraph, t)

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

func TestUpdate(t *testing.T) {
	t.Parallel()
	detectors := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext Context
	}{
		{0, "label truc:",
			Context{currentSituation: "label", currentLabel: "truc"}},
		{1, "label truc(variable=0) : #test parsing",
			Context{currentSituation: "label", currentLabel: "truc"}},
		{2, "jump far # no `:` after jump",
			Context{currentSituation: "jump", currentLabel: "far"}},
	}
	for _, tc := range testCases {
		context := Context{}
		context.update(tc.line, detectors)

		if context.tags != tc.updatedContext.tags {
			t.Errorf("Error in tags:\n got %+v\nwant %+v", context.tags, tc.updatedContext.tags)

		}
		if context != tc.updatedContext {
			t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, context.String(), tc.updatedContext.String())
		}

	}
}

func BenchmarkUpdate(b *testing.B) {
	detectors := initializeDetectors()

	for i := 0; i < b.N; i++ {
		context := Context{}
		context.update("label truc: #bla", detectors)
	}
}
