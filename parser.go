package main

import (
	"fmt"
	"regexp"
	"strings"
)

type situation string

const (
	attrIgnore string = "IGNORE"
	attrTitle  string = "TITLE"
	attrBreak  string = "BREAK"
)

const (
	situationJump    situation = "jump"
	situationLabel   situation = "label"
	situationPending situation = ""
)

type Tag struct {
	ignore    bool
	title     bool
	breakFlow bool
}

// Context gives information about the state of the current line of the script
type Context struct {
	currentSituation  situation // current line situation : jump or label ?
	currentLabel      string    // current line label. Empty if keyword is `situationPending`
	linkedToLastLabel bool      // follows a label or not ?
	lastLabel         string    // last label encountered. Empty if not linkedToLastLabel
	tags              Tag
	currentFile       string
}

func parseRenPy(text []string) RenpyGraph {
	g := NewGraph()

	context := Context{}

	for _, line := range text {

		context.update(line)

		switch context.currentSituation {
		case situationLabel:
			g.AddNode(context.tags, context.currentLabel)
			if context.linkedToLastLabel {
				g.AddEdge(context.tags, context.lastLabel, context.currentLabel, "label")
			}
		case situationJump:
			g.AddNode(context.tags, context.currentLabel)

			g.AddEdge(context.tags, context.lastLabel, context.currentLabel, "")
		}

	}

	return g
}

func (context *Context) update(line string) {
	line = strings.TrimSpace(line)

	// Initialises context
	context.init(line)

	// Handles tags
	context.handleTags(line)

	// Handles keywords
	if !context.tags.ignore {
		if context.tags.breakFlow {
			context.lastLabel = ""
			context.linkedToLastLabel = false
		}
		if r, _ := regexp.Compile(`^\s*label ([a-zA-Z0-9_-]+)\s*:\s*(?:#.*)?$`); r.MatchString(line) {
			// LABEL
			labelName := r.FindStringSubmatch(line)[1]

			fmt.Println("LABEEEL", labelName)
			context.currentLabel = labelName
			context.currentSituation = situationLabel

		} else if r, _ := regexp.Compile(`^\s*jump ([a-zA-Z0-9_-]+)\s*(?:#.*)?$`); r.MatchString(line) {
			// JUMP

			labelName := r.FindStringSubmatch(line)[1]

			fmt.Println("JUUUUMP", labelName)
			context.currentLabel = labelName
			context.currentSituation = situationJump
			context.linkedToLastLabel = false
		} else if r, _ := regexp.Compile(`^\s*(#.*)?$`); r.MatchString(line) {
			// COMMENTS

		} else if context.lastLabel != "" {
			// USUAL VN
			// a label is available (from before in the file) and we are after a jump that is not followed by comments or a label
			context.linkedToLastLabel = true
		}
	}

}

func (context *Context) init(line string) {
	// Reset all tags
	context.tags = Tag{}

	// If last line was a label, say it was the last label
	// Current value have no meaning now
	// Refer to `.situation`
	if context.currentSituation == situationLabel {
		context.lastLabel = context.currentLabel
		context.linkedToLastLabel = true
	}
	context.currentLabel = ""
	context.currentSituation = situationPending
}

func (context *Context) handleTags(line string) {
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
			}
		}
	}
}
