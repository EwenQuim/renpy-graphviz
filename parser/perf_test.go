package parser

import (
	"testing"
)

func BenchmarkGetContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRenpyContent("../testCases/complex", RenpyGraphOptions{})
	}
}

var graph RenpyGraph

func BenchmarkGraph(b *testing.B) {
	var g RenpyGraph
	renpyLines := GetRenpyContent("../testCases/complex", RenpyGraphOptions{})
	for i := 0; i < b.N; i++ {
		g, _ = Graph(renpyLines, RenpyGraphOptions{})
	}
	graph = g
}
