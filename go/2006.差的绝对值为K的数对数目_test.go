package leetcode

import "testing"

func TestCoundKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		{"case1", 4, args{[]int{1, 2, 2, 1}, 1}},
		{"case2", 0, args{[]int{1, 3}, 3}},
		{"case3", 1, args{[]int{10, 2, 10, 9, 1, 6, 8, 9, 2, 8}, 5}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := countKDifference(tt.given.nums, tt.given.k)
			if actual != tt.expected {
				t.Errorf("countKDifference(%s): expected %v, actual %v", tt.name, tt.expected, actual)
			}
		})
	}
}
