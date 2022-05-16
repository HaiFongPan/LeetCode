package leetcode

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    string
	}{
		{"case1", 1, "nlaebolko"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := maxNumberOfBalloons(tt.given)
			if actual != tt.expected {
				t.Errorf("(%s): expected %v, actual %v", tt.name, tt.expected, actual)
			}

		})
	}
}
