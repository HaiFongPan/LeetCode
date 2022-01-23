package leetcode

/*
 * @lc app=leetcode.cn id=1716 lang=golang
 *
 * [1716] 计算力扣银行的钱
 */
// 1+2+3+4+5+6+7=28
// 2+3+4+5+6+7+8=35
// @lc code=start
func totalMoney(n int) int {
	w := n / 7
	sums := func(start int, limit int) int {
		sum := 0
		for i := start; i <= limit; i++ {
			sum = sum + i
		}
		return sum
	}

	if w == 0 {
		return sums(1, n)
	}

	t := n % 7
	s0 := 28
	sn := s0 + (w-1)*7
	return w*(s0+sn)/2 + sums(w+1, w+t)
}

// @lc code=end
