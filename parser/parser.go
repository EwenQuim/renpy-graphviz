package parser

import (
	"strings"
)

// Graph creates a RenpyGraph from lines of script.
// That's the main function
func Graph(text []string) RenpyGraph {

	g := NewGraph()

	context := Context{}

	analytics := Analytics{}

	detectors := initializeDetectors()

	for _, line := range text {

		context.update(line, detectors)

		switch context.currentSituation {

		case situationLabel:
			analytics.labels++
			g.AddNode(context.tags, context.currentLabel)
			if context.linkedToLastLabel {
				g.AddEdge(context.tags, context.lastLabel, context.currentLabel)
			}

		case situationJump:
			analytics.jumps++
			g.AddNode(context.tags, context.currentLabel)
			g.AddEdge(context.tags, context.lastLabel, context.currentLabel)

		case situationCall:
			analytics.calls++
			g.AddNode(context.tags, context.currentLabel)
			g.AddEdge(context.tags, context.lastLabel, context.currentLabel)
		}

	}

	// Plug analytics into the model
	g.info = analytics
	return g
}

// updates the context according to a line of text and detectors
func (context *Context) update(line string, detect customRegexes) {
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
		if detect.label.MatchString(line) {
			// LABEL
			labelName := detect.label.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationLabel
			if context.linkedToLastLabel {
				context.tags.lowLink = true
			}

		} else if detect.jump.MatchString(line) {
			// JUMP
			labelName := detect.jump.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationJump
			context.linkedToLastLabel = false
		} else if detect.call.MatchString(line) {
			// CALL
			labelName := detect.call.FindStringSubmatch(line)[1]

			context.currentLabel = labelName
			context.currentSituation = situationLabel
			context.linkedToLastLabel = true
			context.tags.callLink = true
		} else if detect.comment.MatchString(line) {
			// COMMENTS

		} else if context.lastLabel != "" {
			// USUAL VN
			// a label is available (from before in the file) and we are after a jump that is not followed by comments or a label
			context.linkedToLastLabel = true
		}
	}

}

func (context *Context) init(line string) {

	// If last line was a label, say it was the last label
	// Current value have no meaning now
	// Refer to `.situation`
	if context.currentSituation == situationLabel {

		if context.tags.gameOver {
			context.currentLabel = context.lastLabel
		} else {
			context.lastLabel = context.currentLabel
		}

		context.linkedToLastLabel = true
	}

	context.currentLabel = ""
	context.currentSituation = situationPending

	// Reset all tags
	context.tags = Tag{}
}
