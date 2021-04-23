package parser

import (
	"strings"
	"testing"
)

func TestAddNode(t *testing.T) {
	t.Parallel()
}

func TestString(t *testing.T) {
	t.Parallel()

	r := strings.NewReplacer(" ", "", "\t", "", "\n", "")

	expected := r.Replace(`
digraph  {

	n3[label="complexe"];
	n2[label="simple ending"];
	n1[label="start"];
	n3->n2[style="dotted"];
	n1->n2;
	n1->n3;
	
}`)
	t.Log(expected)

	renpyLines := GetRenpyContent("../testCases/simple")

	graphResult := Graph(renpyLines)
	result := r.Replace(graphResult.String())

	if result != expected {
		t.Fatalf("unexpected graph: \n%v", result)
	}

}
