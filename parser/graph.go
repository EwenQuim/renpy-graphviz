package parser

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"

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
	nodes map[string]*Node
	graph *dot.Graph
	info  Analytics
}

// NewGraph creates an empty graph
func NewGraph() RenpyGraph {
	g := dot.NewGraph(dot.Directed)
	return RenpyGraph{nodes: make(map[string]*Node), graph: g}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func beautifyLabel(label string, tags Tag) string {
	labelName := label
	re := regexp.MustCompile("[)(_=]")
	labelName = re.ReplaceAllString(label, " ")

	if tags.skipLink {
		return labelName[:len(labelName)-5] + " *"
	} else {
		return labelName
	}
}

// AddNode to the ren'py graph, ignore if label already exists
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
		g.nodes[label].repr.Label(strings.ToUpper(labelName)).Attrs("color", "purple", "style", "bold", "shape", "rectangle")
	} else if tags.gameOver {
		g.nodes[label].repr.Attrs("color", "red", "style", "bold", "shape", "septagon")
	}

}

// AddEdge to the renpy graph
func (g *RenpyGraph) AddEdge(tags Tag, label ...string) {

	parentNode := g.nodes[label[0]]
	childrenNode := g.nodes[label[1]]

	edge := g.graph.Edge(*parentNode.repr, *childrenNode.repr)

	if tags.lowLink {
		edge.Attrs("style", "dotted")
	} else if tags.callLink {
		edge.Attrs("style", "dashed", "color", "red")
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
