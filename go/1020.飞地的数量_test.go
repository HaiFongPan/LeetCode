package leetcode

import "testing"

func TestNumEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		{"case1", 3, args{[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}}},
		{"case2", 0, args{[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}}},
		{"case3", 7, args{[][]int{{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0}, {1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0}, {0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0}, {1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1}}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := numEnclaves(tt.given.grid)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.name, tt.expected, actual)
			}

		})
	}
}
