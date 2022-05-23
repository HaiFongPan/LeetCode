package leetcode

import "testing"

func TestCutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
	}

	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		// name, expectd, args
		{"case1", 6, args{[][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}}},
		{"case2", -1, args{[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}}},
		{"case3", 57, args{[][]int{{54581641, 64080174, 24346381, 69107959}, {86374198, 61363882, 68783324, 79706116}, {668150, 92178815, 89819108, 94701471}, {83920491, 22724204, 46281641, 47531096}, {89078499, 18904913, 25462145, 60813308}}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := cutOffTree(tt.given.forest)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, actual)
			}

		})
	}
}
