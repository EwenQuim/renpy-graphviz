package parser

import "testing"

func TestUpdate(t *testing.T) {
	t.Parallel()
	detectors := initializeDetectors()

	testCases := []struct {
		id             int
		line           string
		updatedContext Context
	}{
		{0, "label truc:",
			Context{currentSituation: "label", currentLabel: "truc"}},
		{1, "label truc(variable=0) : #test parsing",
			Context{currentSituation: "label", currentLabel: "truc"}},
		{2, "jump far # no `:` after jump",
			Context{currentSituation: "jump", currentLabel: "far"}},
	}
	for _, tc := range testCases {
		context := Context{}
		context.update(tc.line, detectors)

		if context.tags != tc.updatedContext.tags {
			t.Errorf("Error in tags:\n got %+v\nwant %+v", context.tags, tc.updatedContext.tags)

		}
		if context != tc.updatedContext {
			t.Errorf("Error in struct %v:\n got %+v\nwant %+v", tc.id, context.String(), tc.updatedContext.String())
		}

	}
}
