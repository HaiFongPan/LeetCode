package leetcode

import "testing"

func TestLongestDiverseString(t *testing.T) {
	type args struct {
		a, b, c int
	}

	var tests = []struct {
		name     string
		given    args
		expected string
	}{
		{"case1", args{7, 1, 0}, "aabaa"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := longestDiverseString(tt.given.a, tt.given.b, tt.given.c)
			if actual != tt.expected {
				t.Errorf("(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}
		})
	}
}
