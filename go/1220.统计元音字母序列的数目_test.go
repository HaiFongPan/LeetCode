package leetcode

import "testing"

func Test_countVowelPermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "n1", args: args{1}, want: 5},
		{name: "n2", args: args{2}, want: 10},
		{name: "n5", args: args{5}, want: 68},
		{name: "n20000", args: args{20000}, want: 759959057},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelPermutation(tt.args.n); got != tt.want {
				t.Errorf("countVowelPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
