package longest_substring_without_repeating

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
		{"empty", "", 0},
		{"dvdf", "dvdf", 3},
		{"anviaj", "anviaj", 5},
		{"tmmzuxt", "tmmzuxt", 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tc.s)
			if got != tc.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
