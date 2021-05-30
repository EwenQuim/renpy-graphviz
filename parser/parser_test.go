package parser

import (
	"fmt"
	"testing"
)

func TestEmptyGraph(t *testing.T) {
	t.Parallel()

	graph := Graph([]string{"nothing"}, RenpyGraphOptions{})

	expectedGraph := RenpyGraph{}

	graph.testGraphEquality(expectedGraph, t)

}

func TestUpdate(t *testing.T) {
	t.Parallel()
	detectors := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		context        Context
		updatedContext Context
	}{
		// ------- Lines follow testing -------
		// Basic label
		{1, "label zero:",
			Context{},
			Context{currentSituation: situationLabel, currentLabel: "zero"},
		},
		// Implicit jump
		{2, "label first:",
			Context{currentSituation: situationLabel, currentLabel: "zero"},
			Context{currentSituation: situationLabel, currentLabel: "first", labelLinkedToLastLabel: true, lastLabel: "zero", tags: Tag{lowLink: true}},
		},
		// Context update as there is nothing
		{3, "useless line",
			Context{currentSituation: situationLabel, currentLabel: "first", labelLinkedToLastLabel: true, lastLabel: "zero", tags: Tag{lowLink: true}},
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
		},
		// Context transfer as there is nothing
		{4, "useless line again",
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
		},
		// Jumps
		{5, "jump second # and now jump!",
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
			Context{currentSituation: situationJump, lastLabel: "first", currentLabel: "second"},
		},
		{6, "useless line again and again",
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
		},
		// Call after the jump
		{7, "call third # and now call !",
			Context{currentSituation: situationPending, lastLabel: "first", labelLinkedToLastLabel: true},
			Context{currentSituation: situationCall, lastLabel: "first", currentLabel: "third", labelLinkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		// Call is now used as a previous label
		{8, "useless line again and again",
			Context{currentSituation: situationCall, lastLabel: "first", currentLabel: "third", labelLinkedToLastLabel: true, tags: Tag{callLink: true}},
			Context{currentSituation: situationPending, lastLabel: "third", labelLinkedToLastLabel: true},
		},
		// Implicit jump after call
		{9, "label truc(variable=0) : #test parsing",
			Context{currentSituation: situationPending, lastLabel: "third", labelLinkedToLastLabel: true},
			Context{currentSituation: situationLabel, currentLabel: "truc", lastLabel: "third", labelLinkedToLastLabel: true, tags: Tag{lowLink: true}},
		},
		// Return statement acts like BREAK
		{10, "return # breaks link",
			Context{currentSituation: situationLabel, currentLabel: "truc", lastLabel: "third", labelLinkedToLastLabel: true, tags: Tag{lowLink: true}},
			Context{lastLabel: "truc"},
		},
		// ------- Independent testing -------
		// Parsing
		{2, "label truc(variable=0) : #test parsing",
			Context{},
			Context{currentSituation: situationLabel, currentLabel: "truc"},
		},
		{2, "jump far # no `:` after jump",
			Context{},
			Context{currentSituation: situationJump, currentLabel: "far"},
		},
		{2, "call scene # towards temporary label",
			Context{},
			Context{currentSituation: situationCall, currentLabel: "scene", labelLinkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		// Handling call with/out tags and args
		{2, "call scene(4) # towards temporary label",
			Context{},
			Context{currentSituation: situationCall, currentLabel: "scene(4)", labelLinkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		{2, "call scene(4) # renpy-graphviz: GAMEOVER",
			Context{lastLabel: "maybe_end"},
			Context{currentSituation: situationCall, lastLabel: "maybe_end", currentLabel: "scene(4)", labelLinkedToLastLabel: true, tags: Tag{callLink: true, gameOver: true}},
		},
	}
	for _, tc := range testCases {
		t.Run("Running test for context.update(line) function", func(t *testing.T) {
			tc.context.update(tc.line, detectors)

			if tc.context.tags != tc.updatedContext.tags {
				t.Errorf("Error in tags:\n got %+v\nwant %+v", tc.context.tags, tc.updatedContext.tags)
			}
			if tc.context != tc.updatedContext {
				t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, tc.context.String(), tc.updatedContext.String())
			}
		})
	}
}

func TestInit(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		id              int
		context         Context
		expectedContext Context
	}{
		// Empty at the beginning and in most cases
		{0, Context{},
			Context{},
		},
		// Just after a label
		{1, Context{currentSituation: situationLabel, currentLabel: "yo"},
			Context{labelLinkedToLastLabel: true, lastLabel: "yo"},
		},
		// Same
		{2, Context{currentSituation: situationLabel, currentLabel: "second", lastLabel: "first"},
			Context{labelLinkedToLastLabel: true, lastLabel: "second"},
		},
		// Call situation similar to labelSituation
		{3, Context{currentSituation: situationCall, currentLabel: "second", lastLabel: "first"},
			Context{labelLinkedToLastLabel: true, lastLabel: "second"},
		},
		// Call situation where we go to an ending
		{4, Context{currentSituation: situationCall, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{labelLinkedToLastLabel: false, lastLabel: "first"},
		},
		// Jump situation where we go to an ending
		{5, Context{currentSituation: situationJump, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{labelLinkedToLastLabel: false, lastLabel: "first"},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			tc.context.init()

			if tc.context != tc.expectedContext {
				t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, tc.context.String(), tc.expectedContext.String())
			}
		})
	}
}
