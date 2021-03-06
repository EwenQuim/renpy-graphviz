package main

import (
	"fmt"
	"regexp"
	"strings"
)

type situation string

const (
	situationJump    situation = "jump"
	situationCall    situation = "call"
	situationLabel   situation = "label"
	situationPending situation = ""
)

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
	fmt.Println("Parsing .rpy files...")
	g := NewGraph()

	context := Context{}

	for _, line := range text {

		context.update(line)

		switch context.currentSituation {

		case situationLabel:
			g.AddNode(context.tags, context.currentLabel)
			if context.linkedToLastLabel {
				g.AddEdge(context.tags, context.lastLabel, context.currentLabel)
			}

		case situationJump, situationCall:
			g.AddNode(context.tags, context.currentLabel)
			g.AddEdge(context.tags, context.lastLabel, context.currentLabel)
		}

	}

	return g
}

func (context *Context) update(line string) {
	line = strings.TrimSpace(line)

	// Initialises context
	context.init(line)

	// Handles tags
	context.HandleTags(line)

	// Handles keywords
	if !context.tags.ignore {
		if context.tags.breakFlow {
			context.lastLabel = ""
			context.linkedToLastLabel = false
		}
		if r, _ := regexp.Compile(`^\s*label ([a-zA-Z0-9_()-]+)\s*:\s*(?:#.*)?$`); r.MatchString(line) {
			// LABEL
			labelName := r.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationLabel
			if context.linkedToLastLabel {
				context.tags.lowLink = true
			}

		} else if r, _ := regexp.Compile(`^\s*jump ([a-zA-Z0-9_()-]+)\s*(?:#.*)?$`); r.MatchString(line) {
			// JUMP
			labelName := r.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationJump
			context.linkedToLastLabel = false
		} else if r, _ := regexp.Compile(`^\s*call ([a-zA-Z0-9_()-]+)\s*(?:#.*)?$`); r.MatchString(line) {
			// CALL
			labelName := r.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationLabel
			context.linkedToLastLabel = true
			context.tags.callLink = true
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
