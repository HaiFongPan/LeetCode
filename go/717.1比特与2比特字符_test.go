package leetcode

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    []int
	}{
		{"case1", true, []int{1, 1, 0}},
		{"case2", false, []int{1, 1, 1, 0}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := isOneBitCharacter(tt.given)
			if actual != tt.expected {
				t.Errorf("isOneBitCharacter(%s): expected %v, actual %v", tt.name, tt.expected, actual)
			}
		})
	}

}
