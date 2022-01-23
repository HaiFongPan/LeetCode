package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
var r = ""

func getPermutation(n int, k int) string {
	r = ""
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	search60(nums, k)
	return r
}

func search60(nums []int, k int) {
	if len(nums) == 1 {
		r = r + strconv.Itoa(nums[0])
		return
	}
	avg := factorial(len(nums) - 1)
	bucket := k / avg
	kk := k % avg
	if kk == 0 {
		bucket = bucket - 1
		kk = avg
	}
	r = r + strconv.Itoa(nums[bucket])
	nums = append(nums[:bucket], nums[bucket+1:]...)
	search60(nums, kk)
}

func factorial(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r = r * i
	}
	return r
}

// @lc code=end
