package parser

import (
	"fmt"
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
		{1, "truc3map", Tag{}, "truc 3 map"},
		{2, "truc_map", Tag{}, "truc map"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running test %v", tc.id), func(t *testing.T) {

			if beautifyLabel(tc.line, Tag{}) != tc.result {
				t.Errorf("Error in test %v", tc.id)
			}
		})
	}
}
