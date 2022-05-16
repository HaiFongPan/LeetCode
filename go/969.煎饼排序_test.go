package leetcode

import (
	"testing"
)

func TestPancakeSort(t *testing.T) {
	var tests = []struct {
		name     string
		expected []int
		given    []int
	}{
		{"case1", []int{}, []int{3, 2, 4, 1}},
		{"case1", []int{}, []int{1, 2, 3}},
		{"case1", []int{}, []int{4, 5, 2, 6, 7, 1, 8, 3}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := pancakeSort(tt.given)
			if len(actual) != 10 {
				t.Errorf("error")
			}
		})
	}
}
