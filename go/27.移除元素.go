package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] != val {
			left = left + 1
		} else {
			for right >= 0 && nums[right] == val {
				right = right - 1
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
	}
	return left
}

// @lc code=end
