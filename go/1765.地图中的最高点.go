package leetcode

/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start

var dist [4][2]int = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	type height struct {
		i, j int
	}

	q := []height{}
	// init q
	for i, row := range isWater {
		for j, h := range row {
			if h == 1 {
				q = append(q, height{i, j})
				isWater[i][j] = 0
			} else {
				isWater[i][j] = -1
			}
		}
	}
	for len(q) > 0 {
		// pop
		p := q[0]
		q = q[1:]
		for _, d := range dist {
			ni := p.i + d[0]
			nj := p.j + d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && isWater[ni][nj] == -1 {
				isWater[ni][nj] = isWater[p.i][p.j] + 1
				q = append(q, height{ni, nj})
			}
		}
	}

	return isWater
}

// @lc code=end
