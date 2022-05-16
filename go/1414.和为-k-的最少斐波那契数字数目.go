package leetcode

/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] 和为 K 的最少斐波那契数字数目
 */

// @lc code=start
var fibs []int

func findMinFibonacciNumbers(k int) int {
	buildFibs()
	return findN(k, len(fibs)-1)
}

func findN(k, max int) int {
	if k == 0 {
		return 0
	}
	temp := max
	for fibs[temp] > k {
		temp--
	}
	return 1 + findN(k-fibs[temp], temp)
}

func buildFibs() {
	if len(fibs) == 0 {
		fibs = append(fibs, 1)
		fibs = append(fibs, 1)
		for i := 2; i < 44; i++ {
			fibs = append(fibs, fibs[i-1]+fibs[i-2])
		}
	}
}

// @lc code=end
