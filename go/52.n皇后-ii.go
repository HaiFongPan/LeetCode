package leetcode

// @lc code=start
var resultII int
var rowsII []int

func totalNQueens(n int) int {
	rowsII = make([]int, n)
	resultII = 0
	queen2(0)
	return resultII
}

func queen2(row int) {
	n := len(rowsII)
	if row == len(rowsII) {
		resultII = resultII + 1
	}

	for column := 0; column < n; column++ {
		if isOk2(row, column) {
			rowsII[row] = column
			queen2(row + 1)
			rowsII[row] = 0
		}
	}
}

func isOk2(row, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	n := len(rowsII)

	for i := row - 1; i >= 0; i-- {
		if rowsII[i] == column {
			return false
		}
		if leftUp >= 0 && rowsII[i] == leftUp {
			return false
		}
		if rightUp < n && rowsII[i] == rightUp {
			return false
		}

		leftUp = leftUp - 1
		rightUp = rightUp + 1
	}
	return true
}

// @lc code=end

// @lc code=end
