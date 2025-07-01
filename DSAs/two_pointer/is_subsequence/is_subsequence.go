package is_subsequence

// IsSubsequence returns true if s is a subsequence of t.
func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sIndex := 0
	for _, char := range t {
		if sIndex < len(s) && rune(s[sIndex]) == char {
			sIndex++
			if sIndex == len(s) {
				return true
			}
		}
	}
	return false
}
