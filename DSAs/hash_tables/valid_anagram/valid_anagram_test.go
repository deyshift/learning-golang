package is_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "a", true},
		{"a", "aa", false},
		{"", "", true},
		{"abc", "cba", true},
		{"ab", "a", false},
		{"aabbcc", "abcabc", true},
		{"aabbcc", "aabbc", false},
	}
	for _, tc := range tests {
		t.Run(tc.s+"|"+tc.t, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
