package leetcode

import (
	"sort"
)

func cutOffTree(forest [][]int) (res int) {
	var direction [4][2]int = [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}

	type Point struct {
		i, j int
	}

	type Move675 struct {
		p     Point
		steps int
	}

	sorts := []int{}

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if forest[i][j] > 1 {
				sorts = append(sorts, forest[i][j])
			}
		}
	}

	sort.Slice(sorts, func(i, j int) bool { return sorts[i] < sorts[j] })

	var search675 func(i, j, target int) Move675

	search675 = func(i, j, target int) Move675 {
		p := Point{i, j}
		vst := map[Point]bool{p: true}
		queue := []Move675{{p, 0}}
		for len(queue) > 0 {
			q := queue[0]
			if target == forest[q.p.i][q.p.j] {
				return Move675{q.p, q.steps}
			}
			queue = queue[1:]
			for _, d := range direction {
				in := q.p.i + d[0]
				jn := q.p.j + d[1]
				np := Point{in, jn}
				if !outOfBord(in, jn, len(forest), len(forest[0])) && forest[in][jn] != 0 && !vst[np] {
					queue = append(queue, Move675{np, q.steps + 1})
					vst[np] = true
				}
			}
		}
		return Move675{Point{}, -1}
	}
	i, j, res := 0, 0, 0
	for _, v := range sorts {
		m := search675(i, j, v)
		if m.steps == -1 {
			return -1
		} else {
			res = res + m.steps
			i = m.p.i
			j = m.p.j
		}
	}

	return res
}

func outOfBord(i, j, li, lj int) bool {
	return i < 0 || i == li || j < 0 || j == lj
}
