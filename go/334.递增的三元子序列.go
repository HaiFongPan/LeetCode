package leetcode

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := 0
	middle := -1

	for i := 0; i < len(nums); i++ {
		if middle != -1 && nums[i] > nums[middle] {
			return true
		}
		if nums[i] > nums[min] {
			middle = i
		} else {
			min = i
		}
	}
	return false
}

// @lc code=end
