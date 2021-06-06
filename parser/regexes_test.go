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
		{10, "	call scene(scene=-1) # towards temporary label", true},
		{11, "	call scene(-1) # towards temporary label", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if detect.call.MatchString(tc.line) != tc.updatedContext {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
}

func TestMenuStatement(t *testing.T) {
	t.Parallel()
	detect := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext bool
	}{
		{0, "menu:", true},
		{1, "	menu :", true},
		{2, "	menu : # comments", true},
		{3, "	eileen \"I want the menu:\" ", false},
		{4, "	 # comments menu:", false},
		{5, "	  menu menu_label:", true},
		{6, "menu (\"jfk\", screen=\"airport\"):", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if detect.menu.MatchString(tc.line) != tc.updatedContext {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
}

func TestReturnStatement(t *testing.T) {
	t.Parallel()
	detect := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext bool
	}{
		{0, "    return", true},
		{1, "return", true},
		{2, "    return # comments", true},
		{3, "        return # comments", false},
		{4, "	eileen \"I want to return:\" ", false},
		{5, "	 # comments return", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if detect.returns.MatchString(tc.line) != tc.updatedContext {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
}

func TestChoiceRegex(t *testing.T) {
	t.Parallel()
	detect := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext bool
		choice         string
	}{
		{0, "    \"choice one\":", true, "choice one"},
		{1, "	'choice one':", true, "choice one"},
		{2, "    \"Ne pas parler\" (\"default\", 8.0):", true, "Ne pas parler"},
		{3, "        \"choice one\": # comments", true, "choice one"},
		{4, "	j \"Exact.\" ", false, ""},
		{5, "	 \"exact\" # comments return", false, ""},
		{6, "	\"It's a videogame.\":", true, "It's a videogame."},
		{7, "	\"\\\"It's a videogame.\\\"\":", true, "\"It's a videogame.\""},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if detect.choice.MatchString(tc.line) != tc.updatedContext {
				t.Errorf("Error in test %v", tc.id)
			} else if tc.updatedContext {
				if choice := detect.getChoice(tc.line); choice != tc.choice {
					// t.Logf("%+v", sub)
					t.Errorf("Error in test %v \n%v != %v", tc.id, choice, tc.choice)
				}
			}
		})
	}
}

func TestGetIndent(t *testing.T) {
	detect := initializeDetectors()
	t.Parallel()

	testCases := []struct {
		line    string
		indents int
	}{
		{"jump truc", 0},
		{"    jump truc", 4},
		{"        jump truc", 8},
		{"      jump truc", 6},
		{" jump truc", 1},
	}

	for _, tc := range testCases {
		t.Run("Detect indentation ", func(t *testing.T) {
			indent := detect.getIndent(tc.line, Tag{})

			if indent != tc.indents {
				t.Errorf("Error in indent test:\n got %+v\nwant %+v", indent, tc.indents)
			}
		})
	}
}
