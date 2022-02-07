package leetcode

/*
 * @lc app=leetcode.cn id=1725 lang=golang
 *
 * [1725] 可以形成最大正方形的矩形数目
 */

// @lc code=start
func countGoodRectangles(rectangles [][]int) int {
	ans := 0
	maxLen := 0
	for _, v := range rectangles {
		k := min1725(v[0], v[1])
		if maxLen < k {
			maxLen = k
			ans = 1
		} else if k == maxLen {
			ans++
		}
	}

	return ans
}

func min1725(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
