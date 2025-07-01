package jump_game

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"basic true", []int{2, 3, 1, 1, 4}, true},
		{"basic false", []int{3, 2, 1, 0, 4}, false},
		{"single element", []int{0}, true},
		{"two elements true", []int{2, 0}, true},
		{"two elements false", []int{0, 1}, false},
		{"all zeros except first", []int{1, 0, 0, 0}, false},
		{"large jumps", []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, true},
		{"minimum jumps", []int{1, 1, 1, 1}, true},
		{"zero in middle blocks", []int{1, 0, 1, 0}, false},
		{"barely reach end", []int{2, 0, 0}, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CanJump(tc.nums)
			if got != tc.want {
				t.Errorf("CanJump(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
