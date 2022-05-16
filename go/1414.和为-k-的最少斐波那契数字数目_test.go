package leetcode

import "testing"

func Test_findMinFibonacciNumbers(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{7}, 2},
		{"case2", args{10}, 2},
		{"case3", args{19}, 3},
		{"case4", args{645157245}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinFibonacciNumbers(tt.args.k); got != tt.want {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
