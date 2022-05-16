package leetcode

import "testing"

func TestPushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}

	var tests = []struct {
		name     string
		expected string
		given    args
	}{
		// name, expectd, args
		{"case1", "RR.L", args{"RR.L"}},
		{"case2", "LL.RR.LLRRLL..", args{".L.R...LR..L.."}},
		{"case3", "LL.RRRR", args{".L.RR.."}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := pushDominoes(tt.given.dominoes)
			if actual != tt.expected {
				t.Errorf("%s: expected %s, actual %s", tt.name, tt.expected, actual)
			}

		})
	}
}
