package longest_substring_without_repeating

// LengthOfLongestSubstring returns the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
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
