package parser

import "strings"

// We ConsiderAsUseful a file when it is not a translation file,
// or isn't options/gui/screens .rpy
func ConsiderAsUseful(fullPath string) bool {
	switch {
	case strings.Contains(fullPath, "tl/"):
		return false
	case strings.Contains(fullPath, "options.rpy"):
		return false
	case strings.Contains(fullPath, "screens.rpy"):
		return false
	case strings.Contains(fullPath, "gui.rpy"):
		return false
	default:
		return true
	}
}
