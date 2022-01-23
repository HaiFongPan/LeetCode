package leetcode

/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	max := 0
	for i, v := range nums {
		if v > nums[max] {
			max = i
		}
	}

	for i, v := range nums {
		if i == max || v*2 <= nums[max] {
			continue
		}
		return -1
	}
	return max
}

// @lc code=end
