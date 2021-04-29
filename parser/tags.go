package parser

import (
	"regexp"
	"strings"
)

// # renpy-graphviz: GAMEOVER
const (
	attrIgnore   string = "IGNORE"
	attrTitle    string = "TITLE"
	attrBreak    string = "BREAK"
	attrEnding   string = "GAMEOVER"
	attrSkipLink string = "SKIPLINK"
)

// A Tag allows more control on the graph structure
type Tag struct {
	ignore    bool // renpy-graphviz: IGNORE tag
	title     bool // renpy-graphviz: TITLE tag
	breakFlow bool // renpy-graphviz: BREAK tag
	lowLink   bool // style for implicit jumps
	callLink  bool // style for call statement
	gameOver  bool // renpy-graphviz: GAMEOVER tag
	skipLink  bool // renpy-graphviz: SKIPLINK tag
}

var splitCharacters = regexp.MustCompile(`\W+`)

// handleTags detects tags in the given line. See Tag struct
func (context *Context) handleTags(line string) {
	line = strings.ToLower(line)
	if strings.Contains(line, "renpy-graphviz") {
		lineStrings := strings.Split(line, "renpy-graphviz")
		endOfLine := strings.Join(lineStrings[1:], " ") // removes everything before `renpy-graphviz`

		potentialTags := splitCharacters.Split(endOfLine, -1) // separate every word

		for _, tag := range potentialTags { // sorts tags (false is default)
			switch strings.ToUpper(tag) {
			case attrIgnore:
				context.tags.ignore = true
			case attrTitle:
				context.tags.title = true
			case attrBreak:
				context.tags.breakFlow = true
			case attrEnding:
				context.tags.gameOver = true
			case attrSkipLink:
				context.tags.skipLink = true
			}
		}
	}
}
