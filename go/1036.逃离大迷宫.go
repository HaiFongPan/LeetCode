package leetcode

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
	if n < 2 {
		return true
	}
	move := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	blockedMap := make(map[Point]bool)

	for i := 0; i < len(blocked); i++ {
		blockedMap[Point{blocked[i][0], blocked[i][1]}] = true
	}

	var check = func(s, e []int) bool {
		limit := n * (n - 1) / 2
		queue := []Point{{s[0], s[1]}}
		visit := map[Point]bool{{s[0], s[1]}: true}
		for len(queue) > 0 && limit > 0 {
			p := queue[0]
			for _, v := range move {
				xx := p.x + v[0]
				yy := p.y + v[1]
				np := Point{xx, yy}
				if xx >= 0 && xx < 1000000 && yy >= 0 && yy < 1000000 && blockedMap[np] == false && visit[np] == false {
					if xx == e[0] && yy == e[1] {
						return true
					}
					limit = limit - 1
					queue = append(queue, np)
					visit[np] = true
				}
			}
			queue = queue[1:]
		}
		return limit <= 0
	}

	return check(source, target) && check(target, source)
}

type Point struct {
	x int
	y int
}

// @lc code=end
