package leetcode

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
// @lc code=start
func nextPermutation(nums []int) {
	flag := -1
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			flag = i
		}
	}

	if flag != -1 {
		for j := n - 1; j > flag; j-- {
			if nums[j] > nums[flag] {
				nums[j], nums[flag] = nums[flag], nums[j]
				break
			}
		}
	}
	for i, j := flag+1, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end
