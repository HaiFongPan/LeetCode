package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
// @lc code=start
var result [][]string
var rows []int

func solveNQueens(n int) [][]string {
	rows = make([]int, n)
	result = [][]string{}
	queen(0)
	return result
}

func queen(row int) {
	n := len(rows)
	if row == len(rows) {
		currentResult := []string{}
		for i := 0; i < n; i++ {
			s := strings.Repeat(".", rows[i]) + "Q" + strings.Repeat(".", n-rows[i]-1)
			currentResult = append(currentResult, s)
		}
		result = append(result, currentResult)
	}

	for column := 0; column < n; column++ {
		if isOk(row, column) {
			rows[row] = column
			queen(row + 1)
			rows[row] = 0
		}
	}
}

func isOk(row, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	n := len(rows)

	for i := row - 1; i >= 0; i-- {
		if rows[i] == column {
			return false
		}
		if leftUp >= 0 && rows[i] == leftUp {
			return false
		}
		if rightUp < n && rows[i] == rightUp {
			return false
		}

		leftUp = leftUp - 1
		rightUp = rightUp + 1
	}
	return true
}

// @lc code=end
