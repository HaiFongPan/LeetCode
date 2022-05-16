package leetcode

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    []int
	}{
		{"case1", 2, []int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
		{"case2", 10, []int{3, 3, 7, 7, 10, 11, 11}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := singleNonDuplicate(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %d, actual %d", tt.name, tt.expected, actual)
			}

		})
	}

}
