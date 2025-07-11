package best_time_to_buy_and_sell_stock_2

// MaxProfit2 returns the maximum profit with as many transactions as you like.
func MaxProfit2(prices []int) int {
	maxprofit := 0
	for today := 1; today < len(prices); today++ {
		if prices[today] > prices[today-1] {
			maxprofit += prices[today] - prices[today-1]
		}
	}
	return maxprofit
}
