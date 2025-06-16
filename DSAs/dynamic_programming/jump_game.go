//go:build ignore

package main

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		// If current position is beyond our reach, return false
		if i > maxReach {
			return false
		}

		// Update the maximum reachable position
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		// If we can reach or exceed the last index, return true
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return maxReach >= len(nums)-1
}

func init() {
	// Test cases for canJump function
	println("Testing canJump function:")

	// Test case 1: Basic example - can reach end
	nums1 := []int{2, 3, 1, 1, 4}
	result1 := canJump(nums1)
	println("Input: [2,3,1,1,4] -> Can Jump:", result1) // Expected: true

	// Test case 2: Cannot reach end
	nums2 := []int{3, 2, 1, 0, 4}
	result2 := canJump(nums2)
	println("Input: [3,2,1,0,4] -> Can Jump:", result2) // Expected: false

	// Test case 3: Single element
	nums3 := []int{0}
	result3 := canJump(nums3)
	println("Input: [0] -> Can Jump:", result3) // Expected: true (already at last index)

	// Test case 4: Two elements - can jump
	nums4 := []int{2, 0}
	result4 := canJump(nums4)
	println("Input: [2,0] -> Can Jump:", result4) // Expected: true

	// Test case 5: Two elements - cannot jump
	nums5 := []int{0, 1}
	result5 := canJump(nums5)
	println("Input: [0,1] -> Can Jump:", result5) // Expected: false

	// Test case 6: All zeros except first
	nums6 := []int{1, 0, 0, 0}
	result6 := canJump(nums6)
	println("Input: [1,0,0,0] -> Can Jump:", result6) // Expected: false

	// Test case 7: Large jumps
	nums7 := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	result7 := canJump(nums7)
	println("Input: [5,9,3,2,1,0,2,3,3,1,0,0] -> Can Jump:", result7) // Expected: true

	// Test case 8: Minimum jumps needed
	nums8 := []int{1, 1, 1, 1}
	result8 := canJump(nums8)
	println("Input: [1,1,1,1] -> Can Jump:", result8) // Expected: true

	// Test case 9: Zero in middle blocks path
	nums9 := []int{1, 0, 1, 0}
	result9 := canJump(nums9)
	println("Input: [1,0,1,0] -> Can Jump:", result9) // Expected: false

	// Test case 10: Can barely reach end
	nums10 := []int{2, 0, 0}
	result10 := canJump(nums10)
	println("Input: [2,0,0] -> Can Jump:", result10) // Expected: true

	println("All test cases completed!")
}

func main() {
	// The init function will run automatically before main
}
