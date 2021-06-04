package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var ErrorParentNotFound = errors.New("parent label not found")

func DocumentIssue(err error) {
	color.Red("An error occurred trying to make a graph out of your story.")
	fmt.Println(`I am sorry for the inconvenience.
	
To know more about the issue,
- READ THE DETAILS BELOW. 90 percent of bugs can be resolved easily
- Activate debug mode, and good luck (or send it to the developer)
	- From a terminal, run 'renpy-graphviz -debug'
	- OR Set the debug parameter to 'true' in the 'renpy-graphviz.config'

In the meantime, you'll see an incomplete version of your graph displayed.`)
	color.Red(err.Error())

	if errors.Is(err, ErrorParentNotFound) {
		color.Set(color.Bold)
		fmt.Println(`Not parent label/screen were found. This can be:
- an indentation issue:

    label parent:
  jump somehere


- a tag placed at the wrong place

label parent: # renpy-graphviz: IGNORE
    jump somewhere
	
- a typo:

label thing # <-- look, the ':' symbil is missing

- unindented text that do not belong to a label:

label parent:
    "bla blah"
	"bla blah"
"blah blah" # valid Ren'Py but should be indented as the previous lines
jump somewhere`)
		color.Unset()
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to quit")
	reader.ReadString('\n')
}
