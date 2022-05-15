package leetcode

var direction [4][2]int = [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
var max1219 int

func getMaximumGold(grid [][]int) int {
	max1219 = 0
	m := len(grid)
	n := len(grid[0])

	var search1219 func(m Move1219)
	search1219 = func(m Move1219) {
		ii, jj := m.i, m.j
		// check if search end
		if ii < 0 || ii == len(grid) || jj < 0 || jj == len(grid[0]) || grid[ii][jj] == 0 {
			if max1219 < m.val {
				max1219 = m.val
			}
			return
		}
		rec := grid[ii][jj]
		val := m.val + rec
		grid[ii][jj] = 0
		for _, d := range direction {
			in := ii + d[0]
			jn := jj + d[1]
			search1219(Move1219{in, jn, val})
		}
		grid[ii][jj] = rec

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				search1219(Move1219{i, j, 0})
			}
		}
	}

	return max1219
}

type Move1219 struct {
	i, j, val int
}
