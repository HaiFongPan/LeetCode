package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=2045 lang=golang
 *
 * [2045] 到达目的地的第二短时间
 * BFS / 又抄了一题
 */

// @lc code=start
func secondMinimum(n int, edges [][]int, time int, change int) int {
	g := make([][]int, n+1)
	for _, v := range edges {
		i, j := v[0], v[1]
		g[i] = append(g[i], j)
		g[j] = append(g[j], i)
	}

	type timep struct {
		i, dis int
	}

	state := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		state[i] = make([]int, 2)
		state[i][0] = math.MaxInt32
		state[i][1] = math.MaxInt32
	}

	queue := []timep{timep{1, 0}}
	for state[n][1] == math.MaxInt32 {
		p := queue[0]
		queue = queue[1:]
		for _, v := range g[p.i] {
			dist := p.dis + 1
			// BFS 距离是非递减的，所以最短距离求出来之后，下一个能触达的点就是次短
			if dist < state[v][0] {
				state[v][0] = dist
				queue = append(queue, timep{i: v, dis: dist})
			} else if dist < state[v][1] && dist > state[v][0] {
				state[v][1] = dist
				queue = append(queue, timep{i: v, dis: dist})
			}
		}
	}
	// 红灯时间计算的方式
	ans := 0
	for i := 0; i < state[n][1]; i++ {
		if ans%(2*change) >= change {
			ans = ans + 2*change - ans%(2*change)
		}
		ans = ans + time
	}
	return ans

}
