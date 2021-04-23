package parser

import (
	"strings"
	"testing"
)

func TestGetRenpyContent(t *testing.T) {
	result := GetRenpyContent("../testCases")

	expected := strings.Split(`label start:
    "Is my VN simple ?"
    menu:
        "yes it is":
            jump simple_ending
        "no it isn't":
            jump complexe

label complexe:
    eileen "There are a lot of unexpected things, but at the end..."

label simple_ending: # label complex virtually "jumps" into simple_ending
    eileen "The END !"`, "\n")
	for i, line := range expected {
		if line != result[i] {
			t.Fatalf("expected  line:%v\nreceived line:%v", line, result[i])

		}
	}
}
