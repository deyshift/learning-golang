package is_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChars := make(map[rune]int)
	tChars := make(map[rune]int)

	for _, char := range s {
		sChars[char]++
	}

	for _, char := range t {
		tChars[char]++
	}

	for key, count := range sChars {
		if tChars[key] != count {
			return false
		}
	}

	return true
}
