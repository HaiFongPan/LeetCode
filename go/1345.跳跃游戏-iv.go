package leetcode

/*
 * @lc app=leetcode.cn id=1345 lang=golang
 *
 * [1345] 跳跃游戏 IV
 */

// @lc code=start
func minJumps(arr []int) int {
	n := len(arr)
	repeats := map[int][]int{}
	for i := 0; i < len(arr); i++ {
		repeats[arr[i]] = append(repeats[arr[i]], i)
	}
	queue := []pair{{}}
	vst := map[int]bool{0: true}

	for {
		p := queue[0]
		queue = queue[1:]
		index, step := p.index, p.step

		if index == n-1 {
			return step
		}

		for _, v := range repeats[arr[index]] {
			if !vst[v] {
				vst[v] = true
				queue = append(queue, pair{v, step + 1})
			}
		}
		delete(repeats, arr[index])

		if !vst[index+1] {
			vst[index+1] = true
			queue = append(queue, pair{index + 1, step + 1})
		}
		if index >= 1 && !vst[index-1] {
			vst[index-1] = true
			queue = append(queue, pair{index - 1, step + 1})
		}
	}
}

type pair struct {
	index, step int
}

// @lc code=end
