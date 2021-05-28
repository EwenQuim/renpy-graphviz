package parser

import (
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestExampleGraphs(t *testing.T) {

	testCases := []struct {
		pathToScript  string
		options       RenpyGraphOptions
		expectedGraph string
	}{
		{"../testCases/tagsInGame", RenpyGraphOptions{Silent: true, ShowAtoms: true}, `
digraph  {

	n5[label="ending"];
	n6[color="purple",fontsize="16",label="FIRST",shape="rectangle",style="bold"];
	n4[label="indirect label"];
	n8[label="not indirect"];
	n2[label="option one"];
	n3[label="option two"];
	n7[label="second"];
	n1[label="start"];
	n6->n7[style="dotted"];
	n4->n3;
	n2->n4[style="dotted"];
	n3->n5;
	n1->n2;
	n1->n3;

}`},
		{"../testCases/tagsFake", RenpyGraphOptions{Silent: true, ShowAtoms: true}, `
digraph  {

	n1[label="a"];
	n2[label="b"];
	n3[label="c"];
	n4[label="d"];
	n5[label="e"];
	n6[label="f"];
	n9[label="fake one"];
	n11[label="fake two"];
	n12[label="fake two destination"];
	n7[label="g"];
	n8[label="real one"];
	n10[label="real two"];
	n2->n3;
	n4->n6;
	n6->n7;
	n11->n12;
	n7->n5;
	n8->n10[style="dotted"];

}`},
		{"../testCases/simple", RenpyGraphOptions{Silent: true}, `
digraph  {

	n3[label="complexe"];
	n2[label="simple ending"];
	n1[label="start"];
	n3->n2[style="dotted"];
	n1->n2;
	n1->n3;
	
}`},
		{"../testCases/complex", RenpyGraphOptions{Silent: true}, `
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
	
}`},
		{"../testCases/BOM", RenpyGraphOptions{Silent: true}, `
digraph  {

	n1[label="a"];
	n2[label="b"];
	n1->n2;

}`},
		{"../testCases/jumpAfterReturn", RenpyGraphOptions{Silent: true}, `
digraph  {

	n1[label="a"];
	n2[label="b"];
	n1->n2;

}`},
	}
	dmp := diffmatchpatch.New()

	r := strings.NewReplacer(" ", "", "\t", "", "\n", "")
	for _, tc := range testCases {
		t.Run(tc.pathToScript, func(t *testing.T) {

			renpyLines := GetRenpyContent(tc.pathToScript)
			graphResult := Graph(renpyLines, tc.options)
			result := graphResult.String()

			if r.Replace(result) != r.Replace(tc.expectedGraph) {
				diffs := dmp.DiffMain(tc.expectedGraph, result, false)

				t.Fatalf(dmp.DiffPrettyText(diffs))
			}
		})
	}

}
