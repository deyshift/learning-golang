//go:build ignore

package main

func isSubsequence(s string, t string) bool {
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

func init() {
	// Test cases for isSubsequence function
	println("Testing isSubsequence function:")

	// Test case 1: Basic subsequence
	println("'ace' in 'abcde' ->", isSubsequence("ace", "abcde")) // Expected: true

	// Test case 2: Not a subsequence
	println("'aec' in 'abcde' ->", isSubsequence("aec", "abcde")) // Expected: false

	// Test case 3: Empty string (always subsequence)
	println("'' in 'abcde' ->", isSubsequence("", "abcde")) // Expected: true

	// Test case 4: Subsequence in empty string
	println("'ace' in '' ->", isSubsequence("ace", "")) // Expected: false

	// Test case 5: Multiple occurrences
	println("'abc' in 'aabbcc' ->", isSubsequence("abc", "aabbcc")) // Expected: true

	// Test case 6: Characters exist but wrong order
	println("'axc' in 'ahbgdc' ->", isSubsequence("axc", "ahbgdc")) // Expected: false

	println("All isSubsequence tests completed!")
}

func main() {
	// The init function will run automatically before main
}
