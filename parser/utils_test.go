package parser

import "fmt"

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
