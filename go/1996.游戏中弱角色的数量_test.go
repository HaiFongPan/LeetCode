package leetcode

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{properties: [][]int{{5, 5}, {6, 3}, {3, 6}}}, want: 0},
		{name: "case2", args: args{properties: [][]int{{2, 2}, {3, 3}}}, want: 1},
		{name: "case3", args: args{properties: [][]int{{1, 5}, {10, 4}, {4, 3}}}, want: 1},
		{name: "case3", args: args{properties: [][]int{{6, 3}, {5, 5}, {5, 4}, {5, 2}, {5, 4}, {2, 1}}}, want: 2},
		{name: "case3", args: args{properties: [][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
