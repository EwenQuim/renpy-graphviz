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
		line           string
		context        Context
		updatedContext Context
	}{
		// ------- Lines follow testing -------
		// Basic label
		{"label zero:",
			Context{},
			Context{currentSituation: situationLabel, currentLabel: "zero"},
		},
		// Implicit jump
		{"label first:",
			Context{currentSituation: situationLabel, currentLabel: "zero"},
			Context{currentSituation: situationLabel, currentLabel: "first", linkedToLastLabel: true, lastLabel: "zero", tags: Tag{lowLink: true}},
		},
		// Context update as there is nothing
		{"useless line",
			Context{currentSituation: situationLabel, currentLabel: "first", linkedToLastLabel: true, lastLabel: "zero", tags: Tag{lowLink: true}},
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
		},
		// Context transfer as there is nothing
		{"useless line again",
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
		},
		// Jumps
		{"jump second # and now jump!",
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
			Context{currentSituation: situationJump, lastLabel: "first", currentLabel: "second"},
		},
		{"useless line again and again",
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
		},
		// Call after the jump
		{"call third # and now call !",
			Context{currentSituation: situationPending, lastLabel: "first", linkedToLastLabel: true},
			Context{currentSituation: situationCall, lastLabel: "first", currentLabel: "third", linkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		// Call is now used as a previous label
		{"useless line again and again",
			Context{currentSituation: situationCall, lastLabel: "first", currentLabel: "third", linkedToLastLabel: true, tags: Tag{callLink: true}},
			Context{currentSituation: situationPending, lastLabel: "third", linkedToLastLabel: true},
		},
		// Implicit jump after call
		{"label truc(variable=0) : #test parsing",
			Context{},
			Context{currentSituation: situationLabel, currentLabel: "truc"},
		},
		// ------- Independant testing -------
		// Parsing
		{"label truc(variable=0) : #test parsing",
			Context{},
			Context{currentSituation: situationLabel, currentLabel: "truc"},
		},
		{"jump far # no `:` after jump",
			Context{},
			Context{currentSituation: situationJump, currentLabel: "far"},
		},
		{"	call scene # towards temporary label",
			Context{},
			Context{currentSituation: situationCall, currentLabel: "scene", linkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		// Handling call with/out tags and args
		{"	call scene(4) # towards temporary label",
			Context{},
			Context{currentSituation: situationCall, currentLabel: "scene(4)", linkedToLastLabel: true, tags: Tag{callLink: true}},
		},
		{"	call scene(4) # renpy-graphviz: GAMEOVER",
			Context{lastLabel: "maybe_end"},
			Context{currentSituation: situationCall, lastLabel: "maybe_end", currentLabel: "scene(4)", linkedToLastLabel: true, tags: Tag{callLink: true, gameOver: true}},
		},
	}
	for _, tc := range testCases {
		t.Run("Running test for context.update(line) function", func(t *testing.T) {
			tc.context.update(tc.line, detectors)

			if tc.context.tags != tc.updatedContext.tags {
				t.Errorf("Error in tags:\n got %+v\nwant %+v", tc.context.tags, tc.updatedContext.tags)
			}
			if tc.context != tc.updatedContext {
				t.Errorf("Error in struct:\n got %+v\nwant %+v", tc.context.String(), tc.updatedContext.String())
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
			Context{linkedToLastLabel: true, lastLabel: "yo"},
		},
		// Same
		{2, Context{currentSituation: situationLabel, currentLabel: "second", lastLabel: "first"},
			Context{linkedToLastLabel: true, lastLabel: "second"},
		},
		// Call situation similar to labelSituation
		{3, Context{currentSituation: situationCall, currentLabel: "second", lastLabel: "first"},
			Context{linkedToLastLabel: true, lastLabel: "second"},
		},
		// Call situation where we go to an ending
		{4, Context{currentSituation: situationCall, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{linkedToLastLabel: false, lastLabel: "first"},
		},
		// Jump situation where we go to an ending
		{5, Context{currentSituation: situationJump, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{linkedToLastLabel: false, lastLabel: "first"},
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
