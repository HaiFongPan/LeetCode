package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}

// @lc code=end
