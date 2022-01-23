package leetcode

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = true
	}
	return false
}

// @lc code=end
