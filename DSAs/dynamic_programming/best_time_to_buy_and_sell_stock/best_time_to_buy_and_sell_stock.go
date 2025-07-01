package best_time_to_buy_and_sell_stock

// MaxProfit returns the maximum profit from a single buy and sell.
func MaxProfit(prices []int) int {
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
			maxProfit = currentProfit
		}
	}
	return maxProfit
}
