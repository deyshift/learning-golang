package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// If current element is different from previous unique element
		if nums[fast] != nums[slow] {
			slow++                  // Move slow pointer to next position
			nums[slow] = nums[fast] // Place unique element at slow position
		}
	}

	return slow + 1
}

func init() {
	// Test cases for removeDuplicates function
	println("Testing removeDuplicates function:")

	// Test case 1
	nums1 := []int{1, 1, 2}
	result1 := removeDuplicates(nums1)
	println("Input: [1,1,2] -> Length:", result1, "Array:", nums1[:result1]) // Expected: Length: 2, Array: [1,2]

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result2 := removeDuplicates(nums2)
	println("Input: [0,0,1,1,1,2,2,3,3,4] -> Length:", result2, "Array:", nums2[:result2]) // Expected: Length: 5, Array: [0,1,2,3,4]

	// Test case 3
	nums3 := []int{1}
	result3 := removeDuplicates(nums3)
	println("Input: [1] -> Length:", result3, "Array:", nums3[:result3]) // Expected: Length: 1, Array: [1]

	// Test case 4
	nums4 := []int{1, 2, 3, 4, 5}
	result4 := removeDuplicates(nums4)
	println("Input: [1,2,3,4,5] -> Length:", result4, "Array:", nums4[:result4]) // Expected: Length: 5, Array: [1,2,3,4,5]

	// Test case 5
	nums5 := []int{1, 1, 1, 1, 1}
	result5 := removeDuplicates(nums5)
	println("Input: [1,1,1,1,1] -> Length:", result5, "Array:", nums5[:result5]) // Expected: Length: 1, Array: [1]
}
