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

// func main() {
// 	// Example Usage
// 	println(isSubsequence("ace", "abcde"))  // Output: true
// 	println(isSubsequence("aec", "abcde"))  // Output: false
// 	println(isSubsequence("", "abcde"))     // Output: true
// 	println(isSubsequence("ace", ""))       // Output: false
// 	println(isSubsequence("abc", "aabbcc")) // Output: true
// 	println(isSubsequence("axc", "ahbgdc")) // Output: false
// }
