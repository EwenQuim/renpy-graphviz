package parser

import (
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/require"
)

func TestExampleGraphs(t *testing.T) {
	testCases := []struct {
		pathToScript  string
		options       RenpyGraphOptions
		expectedGraph string
	}{
		{
			"../testCases/BOM", RenpyGraphOptions{Silent: true}, `
digraph  {

	n1[label="a"];
	n2[label="b"];
	n1->n2;

}`,
		},
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
		{"../testCases/jumpAfterReturn", RenpyGraphOptions{Silent: true}, `
			digraph  {
				
				n1[label="a"];
				n2[label="b"];
				n1->n2;
				
				}`},
		{
			"../testCases/menuChoices", RenpyGraphOptions{Silent: true}, `
					digraph  {
						
						n5[label="five"];
						n4[label="four"];
						n1[label="one"];
						n6[label="six"];
						n3[label="three"];
						n2[label="two"];
						n5->n6[style="dotted"];
						n1->n2;
						n1->n3;
						n1->n4;
						n1->n5;
						n1->n2[style="dotted"];
						
						}`,
		},
		{
			"../testCases/nestedLabels", RenpyGraphOptions{Silent: true}, `
						digraph  {
							
							n1[label="first"];
							n6[label="inception"];
							n7[label="inception 2"];
							n8[label="inception 3"];
							n9[label="otherLabelFromThird"];
							n3[label="second"];
							n4[label="selfLoop"];
							n5[label="third"];
							n2[label="useless should be removed"];
							n1->n2[arrowhead="diamond",style="dotted"];
							n1->n3[style="dotted"];
							n6->n7;
							n6->n7[arrowhead="diamond",style="dotted"];
							n6->n8;
							n7->n6;
							n3->n4[arrowhead="diamond",style="dotted"];
							n3->n5;
							n4->n4;
							n5->n6[arrowhead="diamond",style="dotted"];
							n5->n9;
							
							}`,
		},
		{"../testCases/screens", RenpyGraphOptions{Silent: true, ShowScreens: true, ShowNestedScreens: true}, `
						digraph  {
							
							n7[label="game"];
							n3[label="label continue"];
							n2[color="blue",label="nested screen",shape="egg",style="bold"];
							n6[color="blue",label="new screen",shape="egg",style="bold"];
							n8[label="next"];
							n4[color="blue",label="other screen",shape="egg",style="bold"];
							n5[color="blue",label="screen 2",shape="egg",style="bold"];
							n1[color="blue",label="test",shape="egg",style="bold"];
							n7->n6[color="blue",style="dashed"];
							n7->n8;
							n7->n4[color="blue",style="dashed"];
							n5->n2[arrowhead="diamond",arrowtail="inv",color="blue",style="dotted"];
							n5->n6[color="blue",style="dashed"];
							n1->n2[arrowhead="diamond",arrowtail="inv",color="blue",style="dotted"];
							n1->n3[color="blue",style="dashed"];
							n1->n4[color="blue"];
							
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
		{"../testCases/skiplink", RenpyGraphOptions{Silent: true}, `
digraph  {
	
	n6[label="five"];
	n5[label="four"];
	n1[label="one"];
	n7[label="six"];
	n2[label="six *"];
	n4[label="three"];
	n3[label="two"];
	n6->n7[style="dotted"];
	n5->n6[style="dotted"];
	n1->n2;
	n1->n3[style="dotted"];
	n4->n5[style="dotted"];
	n3->n4[style="dotted"];
	
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
	}
	dmp := diffmatchpatch.New()

	r := strings.NewReplacer(" ", "", "\t", "", "\n", "")
	for _, tc := range testCases {
		t.Run(tc.pathToScript, func(t *testing.T) {
			renpyLines := GetRenpyContent(tc.pathToScript, tc.options)
			graphResult, err := Graph(renpyLines, tc.options)
			require.NoError(t, err)

			result := graphResult.String()
			if r.Replace(result) != r.Replace(tc.expectedGraph) {
				diffs := dmp.DiffMain(tc.expectedGraph, result, false)

				t.Fatalf(dmp.DiffPrettyText(diffs))
			}
		})
	}
}
