package leetcode

import "testing"

func TestSortColors(t *testing.T) {
	type args struct {
		nums []int
	}

	var tests = []struct {
		name     string
		expected []int
		given    args
	}{
		// name, expectd, args
		{"case1", []int{0, 1}, args{[]int{1, 0}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.given.nums)
		})
	}
}
