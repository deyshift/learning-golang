//go:build ignore

package main

func isValid(s string) bool {
	// Map of opening brackets to their corresponding closing brackets
	parens := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// Stack to keep track of opening brackets
	var stack []rune

	for _, char := range s {
		// If it's an opening bracket, push to stack
		if _, exists := parens[char]; exists {
			stack = append(stack, char)
		} else {
			// If it's a closing bracket, check if it matches the top of stack
			if len(stack) == 0 {
				return false // No opening bracket to match
			}

			// Pop from stack and check if it matches
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if parens[top] != char {
				return false // Mismatched brackets
			}
		}
	}

	// Valid if stack is empty (all brackets matched)
	return len(stack) == 0
}

func init() {
	// Test cases for isValid function
	println("Testing isValid function:")
	println("'()' ->", isValid("()"))         // Expected: true
	println("'()[]{}' ->", isValid("()[]{}")) // Expected: true
	println("'(]' ->", isValid("(]"))         // Expected: false
	println("'([)]' ->", isValid("([)]"))     // Expected: false
	println("'{[]}' ->", isValid("{[]}"))     // Expected: true
	println("'' ->", isValid(""))             // Expected: true
	println("'(((' ->", isValid("((("))       // Expected: false
	println("')))' ->", isValid(")))"))       // Expected: false
	println("'({[]})' ->", isValid("({[]})")) // Expected: true
}

func main() {
	// Empty main function - examples run automatically via init()
}
