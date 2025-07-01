package remove_duplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		wantLen  int
		wantHead []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{1}},
	}
	for _, tc := range tests {
		t.Run("input", func(t *testing.T) {
			nums := make([]int, len(tc.input))
			copy(nums, tc.input)
			gotLen := RemoveDuplicates(nums)
			if gotLen != tc.wantLen {
				t.Errorf("RemoveDuplicates(%v) = %v, want %v", tc.input, gotLen, tc.wantLen)
			}
			for i := 0; i < gotLen; i++ {
				if nums[i] != tc.wantHead[i] {
					t.Errorf("After RemoveDuplicates(%v), nums[%d]=%d, want %d", tc.input, i, nums[i], tc.wantHead[i])
				}
			}
		})
	}
}
