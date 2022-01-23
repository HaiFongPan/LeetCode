package leetcode

import "testing"

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "n3k3", args: args{n: 3, k: 3}, want: "213"},
		{name: "n4k9", args: args{n: 4, k: 9}, want: "2314"},
		{name: "n3k1", args: args{n: 3, k: 1}, want: "123"},
		{name: "n1k1", args: args{n: 1, k: 1}, want: "1"},
		{name: "n9k1111", args: args{n: 9, k: 1111}, want: "123457869"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
