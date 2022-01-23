package leetcode

/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] 删除回文子序列
 */

// @lc code=start
func removePalindromeSub(s string) int {
	for i, j := 0, len(s)-1; i != j && i < j; i++ {
		if s[i] != s[j] {
			return 2
		}
		j = j - 1
	}
	return 1
}

// @lc code=end
