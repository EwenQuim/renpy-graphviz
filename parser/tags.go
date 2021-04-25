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
	ignore    bool
	title     bool
	breakFlow bool
	lowLink   bool
	callLink  bool
	gameOver  bool
	skipLink  bool
}

func (context *Context) handleTags(line string) {
	line = strings.ToLower(line)
	if strings.Contains(line, "renpy-graphviz") {
		lineStrings := strings.Split(line, "renpy-graphviz")
		endOfLine := strings.Join(lineStrings[1:], " ") // removes everything before `renpy-graphviz`
		splitCharacters := regexp.MustCompile(`\W+`)
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
