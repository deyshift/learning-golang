package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
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

func init() {
	// Example test cases - runs automatically when package loads
	println("Testing isPalindrome function:")
	println("'A man, a plan, a canal: Panama' ->", isPalindrome("A man, a plan, a canal: Panama")) // Expected: true
	println("'race a car' ->", isPalindrome("race a car"))                                         // Expected: false
	println("'Madam' ->", isPalindrome("Madam"))                                                   // Expected: true
	println("'12321' ->", isPalindrome("12321"))                                                   // Expected: true
	println("'hello' ->", isPalindrome("hello"))                                                   // Expected: false
	println("'Was it a car or a cat I saw?' ->", isPalindrome("Was it a car or a cat I saw?"))     // Expected: true
	println("'' (empty) ->", isPalindrome(""))                                                     // Expected: true
	println("'Aa' ->", isPalindrome("Aa"))                                                         // Expected: true
	println("'No 'x' in Nixon' ->", isPalindrome("No 'x' in Nixon"))                               // Expected: true
	println("'.,!@#$%^&*()' ->", isPalindrome(".,!@#$%^&*()"))                                     // Expected: true
}

func main() {
	// Empty main function - examples run automatically via init()
}
