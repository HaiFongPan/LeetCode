package leetcode

import "testing"

func TestGetMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{0, 0}, {0, 0}}}, 0},
		{"case2", args{[][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}}, 24},
		{"case3", args{[][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getMaximumGold(tt.args.grid)
			if actual != tt.want {
				t.Errorf("getMaximumGold() = want %v, actual %v", tt.want, actual)
			}

		})
	}
}
