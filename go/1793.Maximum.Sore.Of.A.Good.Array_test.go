package leetcode

import "testing"

func TestMaximumSore(t *testing.T) {

	type args struct {
		arr []int
		k   int
	}

	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		// name, expectd, args
		{"case1", 15, args{[]int{1, 4, 3, 7, 4, 5}, 3}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := maximumScore(tt.given.arr, tt.given.k)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, actual)
			}

		})
	}
}
