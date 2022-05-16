package leetcode

func luckyNumbers(matrix [][]int) (ans []int) {

	minRow := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		minRow[i] = 0
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				minRow[i] = j
			}
		}
	}

	for i, minj := range minRow {
		isAns := true
		max := matrix[i][minj]
		for ii := 0; ii < len(matrix); ii++ {
			if matrix[ii][minj] > max {
				isAns = false
				break
			}
		}
		if isAns {
			ans = append(ans, matrix[i][minj])
		}
	}
	return ans
}
