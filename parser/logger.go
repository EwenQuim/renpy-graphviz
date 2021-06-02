package parser

import (
	"fmt"

	"github.com/fatih/color"
)

// func (g RenpyGraph) log(a ...interface{}) {
// 	if g.Options.FullDebug {
// 		fmt.Println(a...)
// 	}
// }

func (g RenpyGraph) logLineContext(line string, context, oldContext Context) {
	if g.Options.FullDebug {
		fmt.Println(line)
		if context.diff(oldContext) {
			color.Set(color.Bold)
			fmt.Println("↳", context.String())
			color.Unset()
		}
	}
}
