package jump_game

// CanJump returns true if you can reach the last index.
func CanJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return maxReach >= len(nums)-1
}
