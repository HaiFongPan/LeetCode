package leetcode

/*
 * @lc app=leetcode.cn id=1688 lang=golang
 *
 * [1688] 比赛中的配对次数
 */

// @lc code=start
func numberOfMatches(n int) int {
	times := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			times += n
		} else {
			times += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return times
}

// @lc code=end
