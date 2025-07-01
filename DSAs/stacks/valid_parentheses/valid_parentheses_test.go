package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(((", false},
		{")))", false},
		{"({[]})", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsValid(tc.input)
			if got != tc.want {
				t.Errorf("IsValid(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
