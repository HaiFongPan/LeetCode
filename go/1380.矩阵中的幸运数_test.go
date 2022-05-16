package leetcode

import "testing"

func TestLuckyNumbers(t *testing.T) {
	var tests = []struct {
		name     string
		expected []int
		given    [][]int
	}{
		{"case1", []int{15}, [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := luckyNumbers(tt.given)

			for i, v := range tt.expected {
				if v != actual[i] {
					t.Errorf("luckyNumbers(%s): expected %v, actual %v", tt.name, tt.expected, actual)
					break
				}
			}

		})
	}
}
