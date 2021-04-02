package parser

import (
	"fmt"
	"testing"
)

func TestHandleTags(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		id         int
		line       string
		updatedTag Tag
	}{
		// Basic
		{0, "label truc: # renpy-graphviz: IGNORE", Tag{ignore: true}},
		// Lot of comments
		{1, "label truc(variable='arg') : # bla renpy-graphviz: this is a TITLE", Tag{title: true}},
		// Empty
		{2, "jump truc:", Tag{}},
		// Comments but nothing special
		{3, " narr \"This is a test\" # nothing special in comments", Tag{}},
		// triggers renpy-graphviz but no more
		{4, " narr \"This is a test\" # renpy-graphviz: nothing special", Tag{}},
		// case sensitivity
		{5, " narr \"This is a test\" # renpY-grapHvIz: TITLE", Tag{title: true}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			context := Context{}
			context.handleTags(tc.line)
			if context.tags != tc.updatedTag {
				t.Errorf("Error in tags test %v:\n got %+v\nwant %+v", tc.id, context.tags, tc.updatedTag)

			}
		})

	}
}
