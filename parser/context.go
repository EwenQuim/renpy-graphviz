package parser

import (
	"regexp"
)

// Context gives information about the state of the current line of the script
type Context struct {
	currentSituation  situation // current line situation : jump or label ?
	currentLabel      string    // current line label. Empty if keyword is `situationPending`
	linkedToLastLabel bool      // follows a label or not ?
	lastLabel         string    // last label encountered. Empty if not linkedToLastLabel
	tags              Tag
	// currentFile       string
}

type customRegexes struct {
	label   *regexp.Regexp
	jump    *regexp.Regexp
	call    *regexp.Regexp
	comment *regexp.Regexp
}

func initializeDetectors() customRegexes {
	labelDetector, _ := regexp.Compile(`^\s*label ([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= -]*\))?\s*:\s*(?:#.*)?$`)
	jumpDetector, _ := regexp.Compile(`^\s*jump ([a-zA-Z0-9_]+)\s*(?:#.*)?$`)
	callDetector, _ := regexp.Compile(`^\s*call ([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= -"']*\))?\s*(?:#.*)?$`)
	commentDetector, _ := regexp.Compile(`^\s*(#.*)?$`)
	return customRegexes{
		label:   labelDetector,
		jump:    jumpDetector,
		call:    callDetector,
		comment: commentDetector,
	}
}
