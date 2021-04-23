package parser

import (
	"strings"
	"testing"
)

func TestAddNode(t *testing.T) {
	t.Parallel()
}

func TestStringSimple(t *testing.T) {
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

func TestStringComplex(t *testing.T) {
	t.Parallel()

	r := strings.NewReplacer(" ", "", "\t", "", "\n", "")

	expected := r.Replace(`
digraph  {

	n2[color="red",label="bad ending"];
	n4[color="purple",label="GOOD ENDING"];
	n5[label="route2"];
	n3[label="routeAlternative"];
	n1[color="purple",label="ROUTEONE"];
	n6[color="purple",label="STAAA AA6RT"];
	n5->n2[style="dotted"];
	n3->n4;
	n1->n2;
	n1->n3[style="dotted"];
	n6->n1;
	n6->n5;
	n6->n3;
	
}`)
	t.Log(expected)

	renpyLines := GetRenpyContent("../testCases/complex")

	graphResult := Graph(renpyLines)
	result := r.Replace(graphResult.String())

	if result != expected {
		t.Fatalf("unexpected graph: \n%v", result)
	}
}
