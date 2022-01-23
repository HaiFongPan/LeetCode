package leetcode

import "testing"

func Test_minJumps(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1", args: args{arr: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}}, want: 3},
		{name: "case 2", args: args{arr: []int{7}}, want: 0},
		{name: "case 3", args: args{arr: []int{7, 6, 9, 6, 9, 6, 9, 7}}, want: 1},
		{name: "case 4", args: args{arr: []int{25, -28, -51, 61, -74, -51, -30, 58, 36, 68, -80, -64, 25, -30, -53, 36, -74, 61, -100, -30, -52}}, want: 4},
		{name: "case 5", args: args{arr: []int{7, 7, 2, 7, 7}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJumps(tt.args.arr); got != tt.want {
				t.Errorf("minJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
