package parser

import (
	"fmt"
	"reflect"
	"strings"
)

// ConsiderAsUseful checks if a file is not a translation file,
// or isn't options/gui/screens .rpy
func ConsiderAsUseful(fullPath string) bool {
	switch {
	case strings.Contains(fullPath, "tl/"):
		return false
	case strings.Contains(fullPath, "00"):
		return false
	case strings.Contains(fullPath, "options.rpy"):
		return false
	case strings.Contains(fullPath, "gui.rpy"):
		return false
	default:
		return true
	}
}

// utils for test functions
func (c *Context) String() string {
	str := ""
	if c.detectImplicitJump {
		str += " d"
	}
	str += fmt.Sprint(" stack:", c.labelStack)
	if c.currentSituation != "" {
		str += fmt.Sprint(" situation:", c.currentSituation)
	}
	if c.currentLabel != "" {
		str += fmt.Sprint(" label:", c.currentLabel)
	}
	if c.lastLabel != "" {
		str += fmt.Sprint(" lastLabel:", c.lastLabel)
	}
	if c.currentScreen != "" {
		str += fmt.Sprint(" screen:", c.currentScreen)
	}
	if c.fromScreenToThis != "" {
		str += fmt.Sprint(" screenTo:", c.fromScreenToThis)
	}
	if c.lastChoice != "" {
		str += fmt.Sprint(" choice:", c.lastChoice)
	}
	if c.menuIndent != 0 {
		str += fmt.Sprint(" menuIndent:", c.menuIndent)
	}
	if c.indent != 0 {
		str += fmt.Sprint(" ind:", c.indent)
	}

	return str
}

// diff is used mainly by the logger
func (c *Context) diff(d Context) bool {
	return c.detectImplicitJump != d.detectImplicitJump ||
		!reflect.DeepEqual(c.labelStack, d.labelStack) ||
		c.currentSituation != d.currentSituation ||
		c.currentLabel != d.currentLabel ||
		c.lastLabel != d.lastLabel ||
		c.currentScreen != d.currentScreen ||
		c.fromScreenToThis != d.fromScreenToThis ||
		c.lastChoice != d.lastChoice ||
		c.menuIndent != d.menuIndent
}

func (c *Context) cleanContextAccordingToIndent(line string) {
	// After a menu (indentation before menu indentation) CANNOT NOT STACK for the moment
	if -1 < c.indent && c.indent <= c.menuIndent {
		c.menuIndent = 0
		c.lastChoice = ""
	}

	if len(strings.TrimLeft(line, " ")) >= 4 && strings.TrimLeft(line, " ")[:4] != "menu" { // exception...
		// Updates label stack
		if c.indent >= 0 {
			for i, record := range c.labelStack {
				if c.indent <= record.indent {
					c.lastLabel = c.labelStack[i].labelName
					c.labelStack = c.labelStack[:i]
					break
				}
			}
		}
	}

	// Updates screen situation
	if c.indent == 0 {
		c.currentScreen = ""
	}
}
