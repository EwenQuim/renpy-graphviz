package parser

import (
	"strings"
)

// # renpy-graphviz: GAMEOVER
const (
	attrIgnore      string = "IGNORE"
	attrTitle       string = "TITLE"
	attrBreak       string = "BREAK"
	attrEnding      string = "GAMEOVER"
	attrSkipLink    string = "SKIPLINK"
	attrInGameLabel string = "INGAME_LABEL"
	attrFakeLabel   string = "FAKE_LABEL"
	attrInGameJump  string = "INGAME_JUMP"
	attrFakeJump    string = "FAKE_JUMP"
)

// A Tag allows more control on the graph structure
type Tag struct {
	ignore            bool // renpy-graphviz: IGNORE tag
	title             bool // renpy-graphviz: TITLE tag
	breakFlow         bool // renpy-graphviz: BREAK tag
	lowLink           bool // style for implicit jumps
	nestedLabel       bool // style for nested labels
	callLink          bool // style for call statement
	gameOver          bool // renpy-graphviz: GAMEOVER tag
	skipLink          bool // renpy-graphviz: SKIPLINK tag
	inGameLabel       bool // renpy-graphviz: INGAME_LABEL(label_name) tag
	fakeLabel         bool // renpy-graphviz: FAKE_LABEL(label_name) tag
	inGameJump        bool // renpy-graphviz: INGAME_JUMP(to_label) tag
	fakeJump          bool // renpy-graphviz: FAKE_JUMP(from_label, to_label) tag
	screenToLabel     bool // jump from a screen to a label
	labelToScreen     bool // jump from a label to a screen
	screenToScreen    bool // jump from a screen to another
	useScreenInScreen bool // usage of a screen inside another
	screen            bool // this node is a screen
}

// handleTags detects tags in the given line. See Tag struct
func (context *Context) handleTags(line string, detect customRegexes) {
	line = strings.ToLower(line)
	if strings.Contains(line, "renpy-graphviz") {
		lineStrings := strings.Split(line, "renpy-graphviz")
		endOfLine := strings.Join(lineStrings[1:], " ")                   // removes everything before `renpy-graphviz`
		potentialTags := detect.tags.FindAllStringSubmatch(endOfLine, -1) //splitCharacters.Split(endOfLine, -1) // separate every word
		for _, tagWithSubs := range potentialTags {                       // sorts tags (false is default)
			switch strings.ToUpper(tagWithSubs[1]) {
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
			case attrInGameLabel:
				context.tags.inGameLabel = true
				context.tagLabel = tagWithSubs[2]
			case attrInGameJump:
				context.tags.inGameJump = true
				context.tagJump = tagWithSubs[2]
			case attrFakeLabel:
				context.tags.fakeLabel = true
				context.tagLabel = tagWithSubs[2]
			case attrFakeJump:
				context.tags.fakeJump = true
				context.tagLabel = tagWithSubs[2]
				context.tagJump = tagWithSubs[3]
			}
		}
	}
}
