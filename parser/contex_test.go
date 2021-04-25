package parser

import (
	"fmt"
	"testing"
)

func TestCallStatement(t *testing.T) {
	t.Parallel()
	detect := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext bool
	}{
		{0, "label truc:", false},
		{1, " label truc(variable=0) : #test parsing", false},
		{2, "	label truc(variable='string var') : #test parsing", false},
		{3, "label truc(variable=\"string var\") : #test parsing", false},
		{4, "jump far # no `:` after jump", false},
		{5, "call scene", true},
		{6, "	call scene # towards temporary label", true},
		{7, "	call scene(scene=1) # towards temporary label", true},
		{8, "	call scene(scene=\"string arg\") # towards temporary label", true},
		{9, "	call scene(scene='string arg') # towards temporary label", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if detect.call.MatchString(tc.line) != tc.updatedContext {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
}
