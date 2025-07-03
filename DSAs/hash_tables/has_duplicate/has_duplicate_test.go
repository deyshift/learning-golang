package has_duplicate

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{2, 2, 2, 2}, true},
	}
	for _, tt := range tests {
		if got := HasDuplicate(tt.nums); got != tt.expected {
			t.Errorf("HasDuplicate(%v) = %v; want %v", tt.nums, got, tt.expected)
		}
	}
}
