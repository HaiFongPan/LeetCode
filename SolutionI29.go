package leetcode

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	r := []int{}
	if row == 0 {
		return r
	}
	col := len(matrix[0])
	add := true
	i, j := 0, -1
	for row > 0 && col > 0 {
		for x := 0; x < col; x++ {
			if add {
				j++
			} else {
				j--
			}
			r = append(r, matrix[i][j])
		}

		for y := 0; y < row-1; y++ {
			if add {
				i++
			} else {
				i--
			}
			r = append(r, matrix[i][j])
		}
		add = !add
		col--
		row--
	}

	return r
}
