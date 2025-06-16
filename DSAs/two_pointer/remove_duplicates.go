//go:build ignore

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
	print("Input: [1,1,2] -> Length: ", result1, " Array: [")
	for i := 0; i < result1; i++ {
		if i > 0 {
			print(",")
		}
		print(nums1[i])
	}
	println("]") // Expected: Length: 2, Array: [1,2]

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result2 := removeDuplicates(nums2)
	print("Input: [0,0,1,1,1,2,2,3,3,4] -> Length: ", result2, " Array: [")
	for i := 0; i < result2; i++ {
		if i > 0 {
			print(",")
		}
		print(nums2[i])
	}
	println("]") // Expected: Length: 5, Array: [0,1,2,3,4]
	// Test case 3
	nums3 := []int{1}
	result3 := removeDuplicates(nums3)
	print("Input: [1] -> Length: ", result3, " Array: [")
	for i := 0; i < result3; i++ {
		if i > 0 {
			print(",")
		}
		print(nums3[i])
	}
	println("]") // Expected: Length: 1, Array: [1]

	// Test case 4
	nums4 := []int{1, 2, 3, 4, 5}
	result4 := removeDuplicates(nums4)
	print("Input: [1,2,3,4,5] -> Length: ", result4, " Array: [")
	for i := 0; i < result4; i++ {
		if i > 0 {
			print(",")
		}
		print(nums4[i])
	}
	println("]") // Expected: Length: 5, Array: [1,2,3,4,5]

	// Test case 5
	nums5 := []int{1, 1, 1, 1, 1}
	result5 := removeDuplicates(nums5)
	print("Input: [1,1,1,1,1] -> Length: ", result5, " Array: [")
	for i := 0; i < result5; i++ {
		if i > 0 {
			print(",")
		}
		print(nums5[i])
	}
	println("]") // Expected: Length: 1, Array: [1]

	println("All removeDuplicates tests completed!")
}

func main() {
	// The init function will run automatically before main
}
