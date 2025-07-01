package valid_palindrome

import (
	"strings"
	"unicode"
)

// IsPalindrome returns true if the string is a valid palindrome (ignoring non-alphanumeric and case).
func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
