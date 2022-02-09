package leetcode

import "testing"

func TestGridIllumination(t *testing.T) {
	type args struct {
		n       int
		lamps   [][]int
		queries [][]int
	}

	var tests = []struct {
		name     string
		expected []int
		given    args
	}{
		{"case1", []int{1, 0}, args{5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}}}},
		{"case2", []int{1, 1}, args{5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}}}},
		{"case3", []int{1, 1, 0}, args{5, [][]int{{0, 0}, {0, 4}}, [][]int{{0, 4}, {0, 1}, {1, 4}}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := gridIllumination(tt.given.n, tt.given.lamps, tt.given.queries)
			if len(actual) != len(tt.expected) {
				t.Errorf("(%s): expected len %d, actual len %d", tt.name, len(tt.expected), len(actual))
			} else {
				for i := 0; i < len(actual); i++ {
					if actual[i] != tt.expected[i] {
						t.Errorf("(%s): expected %v, actual %v", tt.name, tt.expected, actual)
						break
					}
				}
			}

		})
	}
}
