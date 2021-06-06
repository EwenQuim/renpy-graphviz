package parser

import (
	"regexp"
	"strings"
)

type customRegexes struct {
	screen            *regexp.Regexp
	useScreenInScreen *regexp.Regexp
	screenToScreen    *regexp.Regexp
	screenToLabel     *regexp.Regexp
	labelToScreen     *regexp.Regexp
	label             *regexp.Regexp
	jump              *regexp.Regexp
	call              *regexp.Regexp
	comment           *regexp.Regexp // line with only blank spaces and comments -will be ignored
	returns           *regexp.Regexp
	menu              *regexp.Regexp
	spaces            *regexp.Regexp
	choice            *regexp.Regexp
	tags              *regexp.Regexp
}

// should be called rarely
func initializeDetectors() customRegexes {
	return customRegexes{
		screen:            regexp.MustCompile(`^\s*(?:init\s+[-\w]*\s+)?screen\s+([a-zA-Z0-9._-]+).*\s*:\s*(?:#.*)?$`),
		useScreenInScreen: regexp.MustCompile(`^\s*use\s+(\w*).*(?:#.*)?$`),
		screenToScreen:    regexp.MustCompile(`^\s*action\s+.*Show\("(.*?)".*\).*(?:#.*)?$`),
		screenToLabel:     regexp.MustCompile(`^\s*action\s+.*(?:Jump|Call)\("(.*?)".*\).*(?:#.*)?$`),
		labelToScreen:     regexp.MustCompile(`^\s*call\s+screen (\w*).*(?:#.*)?$`),
		label:             regexp.MustCompile(`^\s*label\s+([a-zA-Z0-9._-]+)(?:\([a-zA-Z0-9,_= \-"']*\))?\s*:\s*(?:#.*)?$`),
		jump:              regexp.MustCompile(`^\s*jump\s+([a-zA-Z0-9_.]+)\s*(?:#.*)?$`),
		call:              regexp.MustCompile(`^\s*call\s+(([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9,_= \-"']*\))?)\s*(?:#.*)?$`),
		comment:           regexp.MustCompile(`^\s*(#.*)?$`),
		returns:           regexp.MustCompile(`^\s{0,4}return\s*(?:#.*)?$`),
		menu:              regexp.MustCompile(`^\s*menu.*:\s*(?:#.*)?$`),
		spaces:            regexp.MustCompile(`^(\s*).*$`),
		choice:            regexp.MustCompile(`^\s*(?:"(.*?[^\\])"|'(.*?[^\\])').*:\s*(?:#.*)?$`),
		tags:              regexp.MustCompile(`(\w+)(?: *\( *(\w+)(?: *, *(\w+))? *\))?`), // https://regex101.com/r/1vvDF1/1
	}
}

// returns -1 if no submath were found
func (c customRegexes) getIndent(line string, tags Tag) int {
	if tags.inGameLabel || tags.inGameJump {
		return tags.inGameIndent
	}
	if c.comment.MatchString(line) {
		return -1
	}
	return len(c.spaces.FindStringSubmatch(line)[1])
}

func (c customRegexes) getChoice(line string) string {
	sub := c.choice.FindStringSubmatch(line)
	if len(sub) <= 1 {
		panic("getChoice(): error in choice parsing")
	}
	// `sub` can be of the form "choice" (index 0) or 'choice' (index 1)
	for _, e := range sub[1:] { // sub[0] is the match, index > 1 are regex groups
		if strings.TrimSpace(e) != "" {
			return strings.ReplaceAll(e, "\\", "")
		}
	}
	panic("getChoice(): error in choice parsing")
}
