package leetcode

import "testing"

func test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}

	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		// name, expectd, args
		{"case1", 6, args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}},
		{"case2", 1, args{[]int{1}}},
		{"case3", 23, args{[]int{5, 4, -1, 7, 8}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := maxSubArray(tt.given.nums)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, actual)
			}

		})
	}
}
