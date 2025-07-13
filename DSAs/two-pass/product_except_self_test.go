package product_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{[]int{0, 1, 2, 3}, []int{6, 0, 0, 0}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{5}, []int{1}},
	}
	for _, tt := range tests {
		got := productExceptSelf(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
