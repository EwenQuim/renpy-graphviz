package parser

import (
	"regexp"
	"strings"
)

type customRegexes struct {
	label   *regexp.Regexp
	jump    *regexp.Regexp
	call    *regexp.Regexp
	comment *regexp.Regexp
	returns *regexp.Regexp
	menu    *regexp.Regexp
	spaces  *regexp.Regexp
	choice  *regexp.Regexp
}

func initializeDetectors() customRegexes {
	label := regexp.MustCompile(`^\s*label ([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= \-"']*\))?\s*:\s*(?:#.*)?$`)
	jump := regexp.MustCompile(`^\s*jump ([a-zA-Z0-9_]+)\s*(?:#.*)?$`)
	call := regexp.MustCompile(`^\s*call (([a-zA-Z0-9_-]+)(?:\([a-zA-Z0-9_= \-"']*\))?)\s*(?:#.*)?$`)
	menu := regexp.MustCompile(`^\s*menu\s*:\s*(?:#.*)?$`)
	choice := regexp.MustCompile(`^\s*(?:"(.*?)"|'(.*?)').*:\s*(?:#.*)?$`)
	returns := regexp.MustCompile(`^\s{0,4}return\s*(?:#.*)?$`)
	spaces := regexp.MustCompile(`^(\s*).*$`)
	comment := regexp.MustCompile(`^\s*(#.*)?$`)
	return customRegexes{
		label,
		jump,
		call,
		comment,
		returns,
		menu,
		spaces,
		choice,
	}
}

// returns -1 if no submath were found
func (c customRegexes) getIndent(line string) int {
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
	for _, e := range sub[1:] {
		if strings.TrimSpace(e) != "" {
			return e
		}
	}
	panic("getChoice(): error in choice parsing")
}
