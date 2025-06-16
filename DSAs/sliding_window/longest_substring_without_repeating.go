//go:build ignore

package main

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLength := 0
	charIndexMap := make(map[rune]int)

	for end, char := range s {
		if index, ok := charIndexMap[char]; ok && index >= start {
			start = index + 1
		}
		charIndexMap[char] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func init() {
	// Test cases for lengthOfLongestSubstring function
	println("Testing lengthOfLongestSubstring function:")

	println("'abcabcbb' -> Length:", lengthOfLongestSubstring("abcabcbb")) // Expected: 3
	println("'bbbbb' -> Length:", lengthOfLongestSubstring("bbbbb"))       // Expected: 1
	println("'pwwkew' -> Length:", lengthOfLongestSubstring("pwwkew"))     // Expected: 3
	println("'' -> Length:", lengthOfLongestSubstring(""))                 // Expected: 0
	println("'dvdf' -> Length:", lengthOfLongestSubstring("dvdf"))         // Expected: 3
	println("'anviaj' -> Length:", lengthOfLongestSubstring("anviaj"))     // Expected: 5
	println("'tmmzuxt' -> Length:", lengthOfLongestSubstring("tmmzuxt"))   // Expected: 5

	println("All lengthOfLongestSubstring tests completed!")
}

func main() {
	// The init function will run automatically before main
}
