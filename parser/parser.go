package parser

import (
	"fmt"
)

// Context gives information about the state of the current line of the script
type Context struct {
	// General
	currentSituation   situation // current line situation
	indent             int       // negative = blank line (with comments or not)
	detectImplicitJump bool
	// currentFile       string
	// TODO: errors to display (tags, overflows...)

	// Labels
	currentLabel string       // current line label
	labelStack   []labelStack // stack of all labels encountered
	lastLabel    string       // last label encountered, not in the stack. Empty if not labelLinkedToLastLabel
	// labelLinkedToLastLabel bool         // Can the current line reach the lastLabel?

	// Screens
	currentScreen    string // current line "location" (screen information). Empty if and only if not inside a screen
	fromScreenToThis string // target a node from a screen

	// Menus
	menuIndent int    // 0 = not inside a menu
	lastChoice string // last choice in a menu

	// Tags
	tags     Tag    // see the Tag struct
	tagLabel string // fake label written in a comment
	tagJump  string // fake jump destination written in a comment
}

type labelStack struct {
	indent    int
	labelName string
}

// Graph creates a RenpyGraph from lines of script.
// That's the main function
func Graph(text []string, options RenpyGraphOptions) (RenpyGraph, error) {

	g := NewGraph(options)

	context := Context{}

	analytics := Analytics{}

	detectors := initializeDetectors()

	for _, line := range text {

		oldContext := context
		context.update(line, detectors)

		g.logLineContext(line, context, oldContext)

		switch context.currentSituation {

		case situationLabel:
			analytics.labels++
			g.AddNode(context.tags, context.currentLabel)
			var description, fromLabel string
			if len(context.labelStack) >= 2 { // Stacked labels
				description = "stacked"
				fromLabel = context.labelStack[len(context.labelStack)-2].labelName
				context.tags.nestedLabel = true
			} else if context.detectImplicitJump { // Implicit jump
				description = "implicit"
				fromLabel = context.lastLabel
				context.tags.lowLink = true
			}

			if description != "" {
				if err := g.AddEdge(context.tags, fromLabel, context.currentLabel, context.lastChoice); err != nil {
					return g, fmt.Errorf("%v\nERROR(unexpected %v jump) context: %v \n%w", line, description, context.String(), err)
				}
			}
			context.detectImplicitJump = true

		case situationJump:
			analytics.jumps++
			g.AddNode(context.tags, context.currentLabel)
			if len(context.labelStack) == 0 {
				return g, fmt.Errorf("%v\nERROR(empty labelStack) context: %v\n%w", line, context.String(), ErrorParentNotFound)
			}
			if err := g.AddEdge(context.tags, context.labelStack[len(context.labelStack)-1].labelName, context.currentLabel, context.lastChoice); err != nil {
				return g, fmt.Errorf("%v\nERROR(unexpected pure jump) context: %v \n%w", line, context.String(), err)
			}

		case situationCall:
			analytics.calls++
			g.AddNode(context.tags, context.currentLabel)
			var description, fromLabel string
			if len(context.labelStack) >= 2 { // Stacked labels
				description = "stacked"
				fromLabel = context.labelStack[len(context.labelStack)-2].labelName
				context.tags.nestedLabel = true
			} else if context.detectImplicitJump { // Implicit jump
				description = "implicit"
				fromLabel = context.lastLabel
				context.tags.lowLink = true
			}
			if err := g.AddEdge(context.tags, fromLabel, context.currentLabel, context.lastChoice); err != nil {
				return g, fmt.Errorf("%v\nERROR(unexpected %v call): To label %v \n%w", line, description, context.currentLabel, err)
			}
			context.detectImplicitJump = true

		case situationFakeLabel:
			g.AddNode(context.tags, context.tagLabel)

		case situationFakeJump:
			g.AddNode(context.tags, context.tagLabel)
			g.AddNode(context.tags, context.tagJump)
			g.AddEdge(context.tags, context.tagLabel, context.tagJump, context.lastChoice)

		case situationScreen:
			analytics.screens++
			g.AddNode(context.tags, context.currentScreen)

		case situationFromScreenToOther:
			analytics.fromScreenToOther++
			g.AddNode(context.tags, context.fromScreenToThis)
			if err := g.AddEdge(context.tags, context.currentScreen, context.fromScreenToThis); err != nil {
				return g, fmt.Errorf("%v\nERROR(unexpected jump from screen): on screen %v\nContext: %v\n%w", line, context.fromScreenToThis, context.String(), err)
			}
		}

	}

	// Plug analytics into the model
	g.info = analytics

	if !g.Options.Silent {
		fmt.Println(g.String())
	}
	return g, nil
}

// updates the context according to a line of text and detectors
func (context *Context) update(line string, detect customRegexes) {
	context.init()

	context.handleTags(line, detect)

	context.indent = detect.getIndent(line)
	context.cleanContextAccordingToIndent(line)

	// Handles keywords
	if !context.tags.ignore {

		switch {

		// BREAK -before COMMENTS cause this can be a tag-only line
		case context.tags.breakFlow || detect.returns.MatchString(line):
			context.lastLabel = ""
			context.detectImplicitJump = false

		// FAKES -before COMMENTS cause this can be a tag-only line
		case context.tags.fakeLabel:
			context.currentSituation = situationFakeLabel
		case context.tags.fakeJump:
			context.currentSituation = situationFakeJump

		// LABEL -before COMMENTS cause this can be a tag-only line
		case context.currentScreen == "" && (detect.label.MatchString(line) || context.tags.inGameLabel):
			var labelName string
			if context.tags.inGameLabel {
				labelName = context.tagLabel
			} else {
				labelName = detect.label.FindStringSubmatch(line)[1]
			}
			context.labelStack = append(context.labelStack, labelStack{context.indent, labelName})
			context.currentLabel = labelName
			context.currentSituation = situationLabel

		// JUMP -before COMMENTS cause this can be a tag-only line
		case detect.jump.MatchString(line) || context.tags.inGameJump || detect.labelToScreen.MatchString(line):
			var labelName string
			if context.tags.inGameJump {
				labelName = context.tagJump
			} else if detect.jump.MatchString(line) {
				labelName = detect.jump.FindStringSubmatch(line)[1]
			} else {
				context.tags.labelToScreen = true
				labelName = detect.labelToScreen.FindStringSubmatch(line)[1]
			}
			if context.tags.skipLink {
				labelName = labelName + randSeq(5)
			}
			context.currentLabel = labelName
			context.currentSituation = situationJump
			context.lastLabel = ""
			context.detectImplicitJump = false

		// COMMENTS
		case detect.comment.MatchString(line):
			context.currentSituation = situationPending
			// do nothing but save some regex evaluations

		// SCREEN
		case detect.screen.MatchString(line):
			screenName := detect.screen.FindStringSubmatch(line)[1]
			context.tags.screen = true
			context.currentScreen = screenName
			context.currentSituation = situationScreen

		// FROM SCREEN TO LABEL/SCREEN/NESTED SCREEN
		case context.currentScreen != "" && (detect.screenToLabel.MatchString(line) || detect.screenToScreen.MatchString(line) || detect.useScreenInScreen.MatchString(line)):
			var pseudoLabelName string
			if detect.screenToLabel.MatchString(line) {
				context.tags.screenToLabel = true
				pseudoLabelName = detect.screenToLabel.FindStringSubmatch(line)[1]
			} else if detect.useScreenInScreen.MatchString(line) {
				context.tags.useScreenInScreen = true
				pseudoLabelName = detect.useScreenInScreen.FindStringSubmatch(line)[1]
			} else {
				context.tags.screenToScreen = true
				pseudoLabelName = detect.screenToScreen.FindStringSubmatch(line)[1]
			}
			if context.tags.skipLink {
				pseudoLabelName = pseudoLabelName + randSeq(5)
			}
			context.currentSituation = situationFromScreenToOther
			context.fromScreenToThis = pseudoLabelName

		// CALL
		case detect.call.MatchString(line):
			labelName := detect.call.FindStringSubmatch(line)[1]
			if context.tags.skipLink {
				labelName = labelName + randSeq(5)
			}
			context.labelStack = append(context.labelStack, labelStack{context.indent, labelName})
			context.currentLabel = labelName
			context.currentSituation = situationCall
			context.tags.callLink = true

		// MENU
		case detect.menu.MatchString(line):
			context.menuIndent = context.indent

		// CHOICE
		case context.menuIndent < context.indent && detect.choice.MatchString(line):
			context.lastChoice = detect.getChoice(line) //detect.choice.FindStringSubmatch(line)[1]

		// USUAL VN - DIALOGUES
		case context.lastLabel != "" || len(context.labelStack) > 0:
			// a label is available (from before in the file) and we are after a jump that is not followed by comments or a label
			context.detectImplicitJump = true

		default:
		}
	}
}

// initialises the context object before reading a new line, with the context of the previous line
func (context *Context) init() {

	context.currentLabel = ""
	context.currentSituation = situationPending
	context.tagLabel = ""
	context.fromScreenToThis = ""
	context.tagJump = ""

	// Reset all tags
	context.tags = Tag{}
}
