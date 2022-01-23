package leetcode

import "testing"

func Test_removePalindromeSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abb", args: args{"abb"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePalindromeSub(tt.args.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
