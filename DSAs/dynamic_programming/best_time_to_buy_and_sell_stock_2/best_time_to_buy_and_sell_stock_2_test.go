package best_time_to_buy_and_sell_stock_2

import "testing"

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"increasing", []int{1, 2, 3, 4, 5}, 4},
		{"decreasing", []int{7, 6, 4, 3, 1}, 0},
		{"single price", []int{5}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxProfit2(tc.prices)
			if got != tc.want {
				t.Errorf("MaxProfit2(%v) = %v, want %v", tc.prices, got, tc.want)
			}
		})
	}
}
