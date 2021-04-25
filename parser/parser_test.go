package parser

import (
	"fmt"
	"testing"
)

func TestEmptyGraph(t *testing.T) {
	t.Parallel()

	graph := Graph([]string{"nothing"})

	expectedGraph := RenpyGraph{}

	graph.testGraphEquality(expectedGraph, t)

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
		{3, "	call scene # towards temporary label",
			Context{currentSituation: "label", currentLabel: "scene", linkedToLastLabel: true, tags: Tag{callLink: true}}},
		{4, "	call scene(4) # towards temporary label",
			Context{currentSituation: "label", currentLabel: "scene", linkedToLastLabel: true, tags: Tag{callLink: true}}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			context := Context{}
			context.update(tc.line, detectors)

			if context.tags != tc.updatedContext.tags {
				t.Errorf("Error in tags:\n got %+v\nwant %+v", context.tags, tc.updatedContext.tags)
			}
			if context != tc.updatedContext {
				t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, context.String(), tc.updatedContext.String())
			}
		})
	}
}

func BenchmarkUpdate(b *testing.B) {
	detectors := initializeDetectors()

	for i := 0; i < b.N; i++ {
		context := Context{}
		context.update("label truc: #bla", detectors)
	}
}
