package leetcode

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	dir := []struct{ x, y int }{{1, 1}, {-1, -1}, {0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, -1}, {0, 0}}
	x := map[int]int{}
	y := map[int]int{}
	neg := map[int]int{}
	pos := map[int]int{}
	lights := map[Point1001]int{}

	ans := make([]int, len(queries))

	for _, lamp := range lamps {
		yi, xi := lamp[0], lamp[1]
		point := Point1001{xi, yi}
		if lights[point] > 0 {
			continue
		}
		x[xi]++      // col
		y[yi]++      // row
		neg[yi+xi]++ // y = t - x
		pos[yi-xi]++ // y = x + t
		lights[point] = 1
	}

	for index, query := range queries {
		//	fmt.Printf("x: %v\ny: %v\nx+y: %v\ny-x: %v\n", x, y, neg, pos)
		//	fmt.Println("==================================")
		yi, xi := query[0], query[1]
		if y[yi] > 0 || x[xi] > 0 || neg[yi+xi] > 0 || pos[yi-xi] > 0 {
			ans[index] = 1
		}
		// extinguish lights
		for _, d := range dir {
			nx := xi + d.x
			ny := yi + d.y
			point := Point1001{nx, ny}
			if lights[point] == 1 {
				lights[point] = 0
				x[nx]--
				y[ny]--
				neg[ny+nx]--
				pos[ny-nx]--
			}
		}
	}

	return ans
}

type Point1001 struct {
	x, y int
}
