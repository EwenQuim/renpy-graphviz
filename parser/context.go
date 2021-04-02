package parser

import "regexp"

func NewContext() Context {
	context := Context{}
	context.detect = initializeDetectors()
	return context
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
