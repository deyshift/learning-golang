//go:build ignore

package main

import "fmt"

func maxProfit2(prices []int) int {
	maxprofit := 0

	for today := 1; today < len(prices); today++ {
		if prices[today] > prices[today-1] {
			maxprofit += prices[today] - prices[today-1]
		}
	}

	return maxprofit
}

func init() {
	// Test case 1: Example from LeetCode [7,1,5,3,6,4]
	prices1 := []int{7, 1, 5, 3, 6, 4}
	result1 := maxProfit2(prices1)
	fmt.Printf("Test 1 - Prices: %v, Max Profit: %d (Expected: 7)\n", prices1, result1)

	// Test case 2: Increasing prices [1,2,3,4,5]
	prices2 := []int{1, 2, 3, 4, 5}
	result2 := maxProfit2(prices2)
	fmt.Printf("Test 2 - Prices: %v, Max Profit: %d (Expected: 4)\n", prices2, result2)

	// Test case 3: Decreasing prices [7,6,4,3,1]
	prices3 := []int{7, 6, 4, 3, 1}
	result3 := maxProfit2(prices3)
	fmt.Printf("Test 3 - Prices: %v, Max Profit: %d (Expected: 0)\n", prices3, result3)

	// Test case 4: Single price [5]
	prices4 := []int{5}
	result4 := maxProfit2(prices4)
	fmt.Printf("Test 4 - Prices: %v, Max Profit: %d (Expected: 0)\n", prices4, result4)

	println("All maxProfit2 tests completed!")
}

func main() {
	// The init function will run automatically before main
}
