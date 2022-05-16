package leetcode

import "testing"

func TestCommonDivisor(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    []int
	}{
		{"case1", 29, []int{319, 377}},
		{"case2", 1, []int{99, 100}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := commonDivisor1447(tt.given[0], tt.given[1])
			if actual != tt.expected {
				t.Errorf("commonDivisor1447(%s): expected %v, actual %v", tt.name, tt.expected, actual)
			}

		})
	}
}

func TestSimplifiedFractions(t *testing.T) {
	var tests = []struct {
		name     string
		expected []string
		given    int
	}{
		{"case1", []string{"1/2"}, 2},
		{"case2", []string{"1/2", "1/3", "2/3"}, 3},
		{"case3", []string{"1/2", "1/3", "1/4", "2/3", "3/4"}, 4},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := simplifiedFractions(tt.given)
			for i := 0; i < len(actual); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("commonDivisor1447(%s): expected %v, actual %v", tt.name, tt.expected, actual)
					break
				}
			}
		})
	}
}
