package leetcode

import "testing"

func TestUniquePath(t *testing.T) {
	type args struct {
		m int
		n int
	}

	var tests = []struct {
		name     string
		expected int
		given    args
	}{
		// name, expectd, args
		{"case1", 28, args{3, 7}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := uniquePaths(tt.given.m, tt.given.n)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, actual)
			}

		})
	}
}
