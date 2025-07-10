package add_two_numbers

import (
	"reflect"
	"testing"
)

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2 []int
		want   []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}}, // 342 + 465 = 807
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}}, // 9999999+9999=10009998
		{[]int{1, 8}, []int{0}, []int{1, 8}},
		{[]int{5}, []int{5}, []int{0, 1}}, // 5+5=10
	}
	for _, tc := range tests {
		l1 := makeList(tc.l1)
		l2 := makeList(tc.l2)
		got := add_two_numbers(l1, l2)
		gotSlice := listToSlice(got)
		if !reflect.DeepEqual(gotSlice, tc.want) {
			t.Errorf("add_two_numbers(%v, %v) = %v, want %v", tc.l1, tc.l2, gotSlice, tc.want)
		}
	}
}
