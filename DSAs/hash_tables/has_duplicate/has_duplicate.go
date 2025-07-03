package has_duplicate

// HasDuplicate returns true if there are any duplicate integers in nums.
func HasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
