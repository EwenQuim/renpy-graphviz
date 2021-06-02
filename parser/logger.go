package parser

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct{}

func (g RenpyGraph) Log(a ...interface{}) {
	if g.Options.FullDebug {
		fmt.Println(a...)
	}
}

func (g RenpyGraph) LogLineContext(line string, context, oldContext Context) {
	if g.Options.FullDebug {
		fmt.Println(line)
		if context.Diff(oldContext) {
			color.Set(color.Bold)
			fmt.Println("↳", context.String())
			color.Unset()
		}
	}
}
