package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyGraph(t *testing.T) {
	t.Parallel()

	graph, err := Graph([]string{"nothing"}, RenpyGraphOptions{Silent: true})
	if err != nil {
		t.Fatal("failed empty graph creation")
	}

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
			Context{currentSituation: situationLabel, currentLabel: "zero", labelStack: []labelStack{{0, "zero"}}},
		},
		// Implicit jump
		{2, "label first:",
			Context{currentSituation: situationLabel, currentLabel: "zero"},
			Context{currentSituation: situationLabel, currentLabel: "first", labelStack: []labelStack{{0, "first"}}},
		},
		// Context update as there is nothing
		{3, "    useless line",
			Context{currentSituation: situationLabel, currentLabel: "first", detectImplicitJump: true, lastLabel: "zero", tags: Tag{lowLink: true}, labelStack: []labelStack{{0, "first"}}},
			Context{currentSituation: situationPending, lastLabel: "zero", detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
		},
		// Context transfer as there is nothing
		{4, "    useless line again",
			Context{currentSituation: situationPending, lastLabel: "first", detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
			Context{currentSituation: situationPending, lastLabel: "first", detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
		},
		// Jumps
		{5, "      jump second # and now jump!",
			Context{currentSituation: situationPending, lastLabel: "first", detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
			Context{currentSituation: situationJump, currentLabel: "second", labelStack: []labelStack{{0, "first"}}, indent: 6},
		},
		{6, "    useless line again and again",
			Context{currentSituation: situationJump, currentLabel: "second", labelStack: []labelStack{{0, "first"}}, indent: 6},
			Context{currentSituation: situationPending, detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
		},
		// Call after the jump
		{7, "    call third # and now call !",
			Context{currentSituation: situationPending, detectImplicitJump: true, labelStack: []labelStack{{0, "first"}}, indent: 4},
			Context{currentSituation: situationCall, currentLabel: "third", labelStack: []labelStack{{0, "first"}, {4, "third"}}, detectImplicitJump: true, tags: Tag{callLink: true}, indent: 4},
		},
		// Call is now used as a previous label
		{8, "    useless line again and again",
			Context{currentSituation: situationCall, currentLabel: "third", labelStack: []labelStack{{0, "first"}, {4, "third"}}, detectImplicitJump: true, tags: Tag{callLink: true}, indent: 4},
			Context{currentSituation: situationPending, lastLabel: "third", labelStack: []labelStack{{0, "first"}}, detectImplicitJump: true, indent: 4},
		},
		// Implicit jump after call
		{9, "label truc(variable=0) : #test parsing",
			Context{currentSituation: situationPending, lastLabel: "third", labelStack: []labelStack{{0, "first"}}, detectImplicitJump: true, indent: 4},
			Context{currentSituation: situationLabel, currentLabel: "truc", lastLabel: "first", labelStack: []labelStack{{0, "truc"}}, detectImplicitJump: true},
		},
		// Return statement acts like BREAK
		{10, "  return # breaks link",
			Context{currentSituation: situationLabel, currentLabel: "truc", lastLabel: "first", labelStack: []labelStack{{0, "truc"}}, detectImplicitJump: true},
			Context{labelStack: []labelStack{{0, "truc"}}, indent: 2},
		},
	}
	for _, tc := range testCases {
		t.Run("Test context.update(line)", func(t *testing.T) {
			tc.context.update(tc.line, detectors)

			if tc.context.tags != tc.updatedContext.tags {
				t.Errorf("Error in tags:\n got %+v\nwant %+v", tc.context.tags, tc.updatedContext.tags)
			}
			if !reflect.DeepEqual(tc.context, tc.updatedContext) {
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
			Context{},
		},
		// Same
		{2, Context{currentSituation: situationLabel, currentLabel: "second", lastLabel: "first"},
			Context{lastLabel: "first"},
		},
		// Call situation similar to labelSituation
		{3, Context{currentSituation: situationCall, currentLabel: "second", lastLabel: "first"},
			Context{lastLabel: "first"},
		},
		// Call situation where we go to an ending
		{4, Context{currentSituation: situationCall, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{lastLabel: "first"},
		},
		// Jump situation where we go to an ending
		{5, Context{currentSituation: situationJump, currentLabel: "ending", lastLabel: "first", tags: Tag{gameOver: true}},
			Context{lastLabel: "first"},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			tc.context.init()

			if tc.context.tags != tc.expectedContext.tags {
				t.Errorf("Error in tags:\n got %+v\nwant %+v", tc.context.tags, tc.expectedContext.tags)
			}
			if !reflect.DeepEqual(tc.context, tc.expectedContext) {
				t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, tc.context.String(), tc.expectedContext.String())
			}
		})
	}
}
