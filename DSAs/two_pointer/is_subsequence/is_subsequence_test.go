package is_subsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"ace", "abcde", true},
		{"aec", "abcde", false},
		{"", "abcde", true},
		{"ace", "", false},
		{"abc", "aabbcc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tc := range tests {
		t.Run(tc.s+"|"+tc.t, func(t *testing.T) {
			got := IsSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsSubsequence(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
