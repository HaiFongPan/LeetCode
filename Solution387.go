package leetcode

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	s := []int{}
	for _, v := range matrix {
		for _, vv := range v {
			s = append(s, vv)
		}
	}
	sort.Ints(s)
	return s[k-1]
}
