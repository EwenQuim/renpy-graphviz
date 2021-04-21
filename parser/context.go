package parser

import (
	"fmt"
	"regexp"
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

type customRegexes struct {
	label   *regexp.Regexp
	jump    *regexp.Regexp
	call    *regexp.Regexp
	comment *regexp.Regexp
}

func initializeDetectors() customRegexes {
	labelDetector, _ := regexp.Compile(`^\s*label ([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= -]*\))?\s*:\s*(?:#.*)?$`)
	jumpDetector, _ := regexp.Compile(`^\s*jump ([a-zA-Z0-9_]+)\s*(?:#.*)?$`)
	callDetector, _ := regexp.Compile(`^\s*call ([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= -]*\))?\s*:\s*(?:#.*)?$`)
	commentDetector, _ := regexp.Compile(`^\s*(#.*)?$`)
	return customRegexes{
		label:   labelDetector,
		jump:    jumpDetector,
		call:    callDetector,
		comment: commentDetector,
	}
}

func (c *Context) String() string {
	str := ""
	if c.currentSituation != "" {
		str += fmt.Sprint(" situation:", c.currentSituation)
	}
	if c.currentLabel != "" {
		str += fmt.Sprint(" label:", c.currentLabel)
	}
	if c.lastLabel != "" {
		str += fmt.Sprint(" last label:", c.lastLabel)
	}
	if c.linkedToLastLabel {
		str += " linked to last label"
	}

	return str
}
