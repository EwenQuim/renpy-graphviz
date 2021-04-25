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
		// Empty
		{0, "jump truc:", Tag{}},
		// Basic : Ignore Label
		{1, "label truc: # renpy-graphviz: IGNORE", Tag{ignore: true}},
		// Lot of comments
		{2, "label truc(variable='arg') : # bla renpy-graphviz: this is a title", Tag{title: true}},
		// Comments but nothing special
		{3, " narr \"This is a test\" # nothing special in comments", Tag{}},
		// triggers renpy-graphviz but no more
		{4, " narr \"This is a test\" # renpy-graphviz: nothing special", Tag{}},
		// case sensitivity
		{5, " narr \"This is a test\" # renpY-grapHvIz: TITLE", Tag{title: true}},
		// flow break
		{6, "  # renpY-grapHvIz: BREAK", Tag{breakFlow: true}},
		// GAME OVER FLOW
		{7, " label truc: # renpY-grapHvIz: GAMEOVER", Tag{gameOver: true}},
		// SKIPLINK
		{8, " label truc: # renpY-grapHvIz: SKIPLINK", Tag{skipLink: true}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {
			t.Parallel()
			context := Context{}
			context.handleTags(tc.line)
			if context.tags != tc.updatedTag {
				t.Errorf("Error in tags test %v:\n got %+v\nwant %+v", tc.id, context.tags, tc.updatedTag)
			}
		})
	}
}
