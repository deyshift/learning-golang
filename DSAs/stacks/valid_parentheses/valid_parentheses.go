package valid_parentheses

// IsValid returns true if the input string is a valid set of parentheses.
func IsValid(s string) bool {
	parens := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune
	for _, char := range s {
		if _, exists := parens[char]; exists {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if parens[top] != char {
				return false
			}
		}
	}
	return len(stack) == 0
}
