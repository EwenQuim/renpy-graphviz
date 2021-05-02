package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestBeautifyLabel(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		id     int
		line   string
		tags   Tag
		result string
	}{
		{0, "truc", Tag{}, "truc"},
		{1, "truc3", Tag{}, "truc 3"},
		{2, "truc3map", Tag{}, "truc 3 map"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if beautifyLabel(tc.line, Tag{}) != tc.result {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
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

	graphResult := Graph(renpyLines, false)
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

	n2[color="red",label="bad ending",shape="septagon",style="bold"];
	n4[color="purple",fontsize="16",label="GOOD ENDING",shape="rectangle",style="bold"];
	n5[label="route 2"];
	n3[label="routeAlternative"];
	n1[color="purple",fontsize="16",label="ROUTEONE",shape="rectangle",style="bold"];
	n6[color="purple",fontsize="16",label="STAAA AA 6 RT",shape="rectangle",style="bold"];
	n5->n2[style="dotted"];
	n3->n4;
	n1->n2;
	n1->n3[style="dotted"];
	n6->n1;
	n6->n5;
	n6->n3;
	
}	`)
	t.Log(expected)

	renpyLines := GetRenpyContent("../testCases/complex")

	graphResult := Graph(renpyLines, false)
	result := r.Replace(graphResult.String())

	if result != expected {
		t.Fatalf("unexpected graph: \n%v", result)
	}
}
