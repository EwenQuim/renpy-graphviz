package parser

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"github.com/emicklei/dot"
)

// A Node is a Ren'Py label and some properties, including its graph representation
type Node struct {
	name      string
	neighbors []string
	repr      *dot.Node
	notAlone  bool
}

// RenpyGraph is the graph of Ren'Py story structure
type RenpyGraph struct {
	nodes   map[string]*Node
	graph   *dot.Graph
	info    Analytics
	Options RenpyGraphOptions
}

// RenpyGraphOptions are options that can be set to make a more customizable graph
type RenpyGraphOptions struct {
	ShowEdgesLabels   bool // Show Labels on Edges? Can be unreadable but there is more information
	ShowAtoms         bool // Show lonely nodes ? Might be useful but useless most of the time - and it avoids writing IGNORE tag everywhere
	ShowNestedScreens bool // Show nested screens (`use` keyword)
	Silent            bool // Display .dot graph in the stdout
	OpenFile          bool // Open the image in the default image viewer or not ?
}

// NewGraph creates an empty graph
func NewGraph(options RenpyGraphOptions) RenpyGraph {
	g := dot.NewGraph(dot.Directed)
	return RenpyGraph{nodes: make(map[string]*Node), graph: g, Options: options}
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var replaceBlanks = regexp.MustCompile("[)(_=]")

func beautifyLabel(label string, tags Tag) string {
	labelName := label

	labelName = replaceBlanks.ReplaceAllString(label, " ")

	if tags.skipLink {
		return labelName[:len(labelName)-5] + " *"
	}

	labelName = insertSpaces(labelName)

	return labelName
}

func insertSpaces(s string) string {
	buf := &bytes.Buffer{}
	var previousChar bool
	var previousNumber bool
	for i, rune := range s {
		nowChar := unicode.IsLetter(rune)
		nowNumber := unicode.IsDigit(rune)
		if previousNumber && nowChar && i > 0 {
			buf.WriteRune(' ')
		} else if previousChar && nowNumber && i > 0 {
			buf.WriteRune(' ')
		}
		previousChar = nowChar
		previousNumber = nowNumber
		buf.WriteRune(rune)
	}
	return buf.String()
}

// AddNode to the ren'py graph. If label already exists, only apply styles
func (g *RenpyGraph) AddNode(tags Tag, label string) {
	if tags.useScreenInScreen && !g.Options.ShowNestedScreens {
		return
	}
	// fmt.Println("adding ", label, "to", g)

	labelName := beautifyLabel(label, tags)

	_, exists := g.nodes[label]
	if !exists {
		nodeGraph := g.graph.Node(label)
		nodeGraph.Label(labelName)

		g.nodes[label] = &Node{name: label, neighbors: make([]string, 0), repr: &nodeGraph}
	}
	if tags.title {
		g.nodes[label].repr.Label(strings.ToUpper(labelName)).Attrs("color", "purple", "style", "bold", "shape", "rectangle", "fontsize", "16")
	} else if tags.gameOver {
		g.nodes[label].repr.Attrs("color", "red", "style", "bold", "shape", "septagon")
	} else if tags.screen {
		g.nodes[label].repr.Attrs("color", "blue", "style", "bold", "shape", "egg")
	}

}

// AddEdge to the renpy graph
func (g *RenpyGraph) AddEdge(tags Tag, label ...string) {
	if tags.useScreenInScreen && !g.Options.ShowNestedScreens {
		return
	}

	parentNode := g.nodes[label[0]]
	childrenNode := g.nodes[label[1]]

	g.nodes[label[0]].notAlone = true
	g.nodes[label[1]].notAlone = true

	edge := g.graph.Edge(*parentNode.repr, *childrenNode.repr)

	if tags.lowLink {
		edge.Attrs("style", "dotted")
	} else if tags.callLink {
		edge.Attrs("style", "dashed", "color", "red")
	} else if tags.screenToLabel || tags.labelToScreen {
		edge.Attrs("style", "dashed", "color", "blue")
	} else if tags.useScreenInScreen {
		edge.Attrs("style", "dotted", "color", "blue", "arrowhead", "diamond", "arrowtail", "inv")
	} else if tags.screenToScreen {
		edge.Attrs("color", "blue")
	}
	if g.Options.ShowEdgesLabels && len(label) >= 3 {
		edge.Label(label[2])
	}

	parentNode.neighbors = append(parentNode.neighbors, label[1])

}

// CreateFile creates a file with the graph description in dot language
// It is meant to be used on a computer
// Calls (renpyGraph).String to output file
func (g *RenpyGraph) CreateFile(fileName string) error {
	b := []byte(g.String())
	return ioutil.WriteFile(fileName, b, 0644)
}

// String returns a string with the graph description in dot language
// It is meant to be used by other libraries or programs
// It removes Atoms if specified in .Options field
func (g *RenpyGraph) String() string {
	g.removeAtomsIfSpecified()
	return g.graph.String()
}

func (g *RenpyGraph) removeAtomsIfSpecified() {
	if !g.Options.ShowAtoms {
		for name, node := range g.nodes {
			if !node.notAlone {
				g.graph.DeleteNode(name)
			}
		}
	}
}
