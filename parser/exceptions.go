package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var ErrorParentNotFound = errors.New("parent label not found")

type ErrorIngameTagIndent struct {
	tagType string // INGAME_LABEL or INGAME_JUMP
	indent  int
	err     error // underlying error, if any
}

func (e ErrorIngameTagIndent) Error() string {
	return fmt.Sprintf("tag %s has an invalid indentation level: %d or an error: %s", e.tagType, e.indent, e.err)
}

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

	color.Set(color.Bold)

	if errors.Is(err, ErrorParentNotFound) {
		fmt.Println(`Not parent label/screen were found. This can be:
- an indentation issue:

    label parent:
  jump somehere


- a tag placed at the wrong place

label parent: # renpy-graphviz: IGNORE
    jump somewhere
	
- a typo:

label thing # <-- look, the ':' symbol is missing

- unindented text that do not belong to a label:

label parent:
    "bla blah"
	"bla blah"
"blah blah" # valid Ren'Py but should be indented as the previous lines
jump somewhere`)
	}
	var ingameTagIndent ErrorIngameTagIndent
	if errors.As(err, &ingameTagIndent) {
		fmt.Println(`The indentation given in the IN_GAME tag is not correct.
INGAME_LABEL and INGAME_JUMP tags need and indentation level
since they behave like real label/jumps`)
	}
	color.Unset()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to quit")
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Couldn't read input: %s\n", err)
	}
}
