package parser

import (
	"bytes"
	"fmt"
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
}

// RenpyGraph is the graph of Ren'Py story structure
type RenpyGraph struct {
	nodes           map[string]*Node
	graph           *dot.Graph
	info            Analytics
	showEdgesLabels bool // Show Labels on Edges? Can be unreadable
}

// NewGraph creates an empty graph
func NewGraph(edges bool) RenpyGraph {
	g := dot.NewGraph(dot.Directed)
	return RenpyGraph{nodes: make(map[string]*Node), graph: g, showEdgesLabels: edges}
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
	}

}

// AddEdge to the renpy graph
func (g *RenpyGraph) AddEdge(tags Tag, label ...string) {

	parentNode := g.nodes[label[0]]
	childrenNode := g.nodes[label[1]]

	fmt.Println(strings.Join(label, ` / `))
	edge := g.graph.Edge(*parentNode.repr, *childrenNode.repr)

	if tags.lowLink {
		edge.Attrs("style", "dotted")
	} else if tags.callLink {
		edge.Attrs("style", "dashed", "color", "red")
	}
	if g.showEdgesLabels && len(label) == 3 {
		edge.Label(label[2])
	}

	parentNode.neighbors = append(parentNode.neighbors, label[1])

}

// CreateFile creates a file with the graph description in dot language
// It is meant to be used on a computer
func (g *RenpyGraph) CreateFile(fileName string) error {
	b := []byte(g.graph.String())
	return ioutil.WriteFile(fileName, b, 0644)
}

// String returns a string with the graph description in dot language
// It is meant to be used by other libraries or programs
func (g *RenpyGraph) String() string {
	return g.graph.String()
}
