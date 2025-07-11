package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"basic example", []int{7, 1, 5, 3, 6, 4}, 5},
		{"decreasing prices", []int{7, 6, 4, 3, 1}, 0},
		{"single day", []int{5}, 0},
		{"two days profit", []int{1, 5}, 4},
		{"two days no profit", []int{5, 1}, 0},
		{"multiple peaks", []int{2, 4, 1, 7, 5, 8}, 7},
		{"same prices", []int{3, 3, 3, 3}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxProfit(tc.prices)
			if got != tc.want {
				t.Errorf("MaxProfit(%v) = %v, want %v", tc.prices, got, tc.want)
			}
		})
	}
}
