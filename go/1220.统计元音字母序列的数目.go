package leetcode

/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] 统计元音字母序列的数目
 */

// @lc code=start
const mod = 1e9 + 7

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 2; j <= n; j++ {
		aj := (e + i + u) % mod
		ej := (a + i) % mod
		ij := (e + o) % mod
		oj := i
		uj := (i + o) % mod

		a, e, i, o, u = aj, ej, ij, oj, uj
	}

	return (a + e + i + o + u) % mod
}

// @lc code=end
