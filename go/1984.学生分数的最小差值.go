package leetcode

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) (ans int) {
	sort.Ints(nums)
	ans = math.MaxInt32
	for i, num := range nums[:len(nums)-k+1] {
		ans = min1984(ans, nums[i+k-1]-num)
	}
	return ans
}

func min1984(x, y int) int {
	if x < y {
		return x
	}
	return y
}
