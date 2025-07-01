package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"Madam", true},
		{"12321", true},
		{"hello", false},
		{"Was it a car or a cat I saw?", true},
		{"", true},
		{"Aa", true},
		{"No 'x' in Nixon", true},
		{".,!@#$%^&*()", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
