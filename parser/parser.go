package parser

import "fmt"

// Context gives information about the state of the current line of the script
type Context struct {
	currentSituation  situation // current line situation : jump or label ?
	currentLabel      string    // current line label. Empty if keyword is `situationPending`
	linkedToLastLabel bool      // follows a label or not ?
	lastLabel         string    // last label encountered. Empty if not linkedToLastLabel
	indent            int       // 4 spaces = 1 indent
	menuIndent        int       // negative = not inside a menu
	lastChoice        string    // last choice in a menu
	tags              Tag       // see the Tag struct
	tagLabel          string    // fake label written in a comment
	tagJump           string    // fake jump destination written in a comment
	// currentFile       string
}

// Graph creates a RenpyGraph from lines of script.
// That's the main function
func Graph(text []string, options RenpyGraphOptions) RenpyGraph {

	g := NewGraph(options)

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
				g.AddEdge(context.tags, context.lastLabel, context.currentLabel, context.lastChoice)
			}

		case situationJump:
			analytics.jumps++
			g.AddNode(context.tags, context.currentLabel)
			g.AddEdge(context.tags, context.lastLabel, context.currentLabel, context.lastChoice)

		case situationCall:
			analytics.calls++
			g.AddNode(context.tags, context.currentLabel)
			if _, exists := g.nodes[context.lastLabel]; !exists {
				println("Error in your game: no label detected before the following line\n", line)
				g.AddNode(Tag{}, context.lastLabel) //Useless but security in case the game isn't well structured
			}
			g.AddEdge(context.tags, context.lastLabel, context.currentLabel, context.lastChoice)

		case situationFakeLabel:
			g.AddNode(context.tags, context.tagLabel)

		case situationFakeJump:
			g.AddNode(context.tags, context.tagLabel)
			g.AddNode(context.tags, context.tagJump)
			g.AddEdge(context.tags, context.tagLabel, context.tagJump, context.lastChoice)
		}

	}

	// Plug analytics into the model
	g.info = analytics

	if !g.Options.Silent {
		fmt.Println(g.String())
	}
	return g
}

// updates the context according to a line of text and detectors
func (context *Context) update(line string, detect customRegexes) {

	context.init()

	context.handleTags(line, detect)

	context.indent = detect.getIndent(line)
	// After a menu (indentation before menu indentation)
	if -1 < context.indent && context.indent <= context.menuIndent {
		context.menuIndent = 0
		context.lastChoice = ""
	}

	// Handles keywords
	if !context.tags.ignore {

		switch {

		// BREAK -before COMMENTS cause this can be a tag-only line
		case context.tags.breakFlow || detect.returns.MatchString(line):
			context.lastLabel = ""
			context.linkedToLastLabel = false

		// FAKES -before COMMENTS cause this can be a tag-only line
		case context.tags.fakeLabel:
			context.currentSituation = situationFakeLabel
		case context.tags.fakeJump:
			context.currentSituation = situationFakeJump

		// LABEL -before COMMENTS cause this can be a tag-only line
		case detect.label.MatchString(line) || context.tags.inGameLabel:
			var labelName string
			if context.tags.inGameLabel {
				labelName = context.tagLabel
			} else {
				labelName = detect.label.FindStringSubmatch(line)[1]
			}

			context.currentLabel = labelName
			context.currentSituation = situationLabel
			if context.linkedToLastLabel {
				context.tags.lowLink = true
			}

		// JUMP -before COMMENTS cause this can be a tag-only line
		case detect.jump.MatchString(line) || context.tags.inGameJump:
			var labelName string
			if context.tags.inGameJump {
				labelName = context.tagJump
			} else {
				labelName = detect.jump.FindStringSubmatch(line)[1]
			}
			if context.tags.skipLink {
				labelName = labelName + randSeq(5)
			}
			context.currentLabel = labelName
			context.currentSituation = situationJump
			context.linkedToLastLabel = false

		// COMMENTS
		case detect.comment.MatchString(line):

		// CALL
		case detect.call.MatchString(line):
			labelName := detect.call.FindStringSubmatch(line)[1]
			if context.tags.skipLink {
				labelName = labelName + randSeq(5)
			}
			context.currentLabel = labelName
			context.currentSituation = situationCall
			context.linkedToLastLabel = true
			context.tags.callLink = true

		// MENU
		case detect.menu.MatchString(line):
			context.menuIndent = context.indent

		// CHOICE
		case context.menuIndent < context.indent && detect.choice.MatchString(line):
			context.lastChoice = detect.getChoice(line) //detect.choice.FindStringSubmatch(line)[1]

		// USUAL VN
		case context.lastLabel != "":
			// a label is available (from before in the file) and we are after a jump that is not followed by comments or a label
			context.linkedToLastLabel = true

		default:
		}
	}
}

// initialises the context object before reading a new line, with the context of the previous line
func (context *Context) init() {

	// If last line was a label, say it was the last label
	// Current value have no meaning now
	// Refer to `.situation`
	if context.currentSituation == situationLabel || context.currentSituation == situationCall {
		// Do not follow "game over" marked tags
		// Keep the previous label if "game over" tag
		// Else, update the corresponding label
		if !context.tags.gameOver {
			context.lastLabel = context.currentLabel
			context.linkedToLastLabel = true
		}
	}

	context.currentLabel = ""
	context.currentSituation = situationPending
	context.tagLabel = ""
	context.tagJump = ""

	// Reset all tags
	context.tags = Tag{}
}
