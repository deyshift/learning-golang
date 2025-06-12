package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		currentProfit := prices[i] - minPrice

		if currentProfit > maxProfit {
			currentProfit = maxProfit
		}
	}

	return maxProfit

}

func init() {
	// Test cases for maxProfit function
	println("Testing maxProfit function:")

	// Test case 1: Basic example
	prices1 := []int{7, 1, 5, 3, 6, 4}
	result1 := maxProfit(prices1)
	println("Input: [7,1,5,3,6,4] -> Max Profit:", result1) // Expected: 5 (buy at 1, sell at 6)

	// Test case 2: Decreasing prices
	prices2 := []int{7, 6, 4, 3, 1}
	result2 := maxProfit(prices2)
	println("Input: [7,6,4,3,1] -> Max Profit:", result2) // Expected: 0 (no profit possible)

	// Test case 3: Single day
	prices3 := []int{5}
	result3 := maxProfit(prices3)
	println("Input: [5] -> Max Profit:", result3) // Expected: 0 (need at least 2 days)

	// Test case 4: Two days - profit possible
	prices4 := []int{1, 5}
	result4 := maxProfit(prices4)
	println("Input: [1,5] -> Max Profit:", result4) // Expected: 4 (buy at 1, sell at 5)

	// Test case 5: Two days - no profit
	prices5 := []int{5, 1}
	result5 := maxProfit(prices5)
	println("Input: [5,1] -> Max Profit:", result5) // Expected: 0 (price decreases)

	// Test case 6: Multiple peaks
	prices6 := []int{2, 4, 1, 7, 5, 8}
	result6 := maxProfit(prices6)
	println("Input: [2,4,1,7,5,8] -> Max Profit:", result6) // Expected: 7 (buy at 1, sell at 8)

	// Test case 7: Same prices
	prices7 := []int{3, 3, 3, 3}
	result7 := maxProfit(prices7)
	println("Input: [3,3,3,3] -> Max Profit:", result7) // Expected: 0 (no price change)
}
