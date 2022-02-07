package leetcode

/*
 * @lc app=leetcode.cn id=2013 lang=golang
 *
 * [2013] 检测正方形
 */

// @lc code=start

type DetectSquares struct {
	// y, x, cnt
	points map[int]map[int]int
}

type Point2013 struct {
	x, y int
}

func Constructor() DetectSquares {
	return DetectSquares{map[int]map[int]int{}}
}

// var direction = [2]int{-1, 1}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if this.points[y] == nil {
		this.points[y] = map[int]int{}
		this.points[y][x] = 1
	} else {
		this.points[y][x] += 1
	}
}

func (this *DetectSquares) Count(point []int) int {
	count := 0
	x, y := point[0], point[1]
	xs := this.points[y]
	if xs == nil {
		return 0
	}

	for i, cnt := range xs {
		dist := abs2013(i, x)
		if dist == 0 {
			continue
		}
		// for _, v := range direction {
		// 	if p2, ok := this.points[y+dist*v][i]; ok {
		// 		if p3, ok := this.points[y+dist*v][x]; ok {
		// 			// fmt.Printf("found points, [%d, %d]=%d, [%d, %d]=%d, [%d, %d]=%d\n",
		// 			// 	i, y, cnt, i, y+dist*v, p2, x, y+dist*v, p3)
		// 			count += cnt * p2 * p3
		// 		}
		// 	}
		// }
		count += cnt * this.points[y+dist][i] * this.points[y+dist][x]
		count += cnt * this.points[y-dist][i] * this.points[y-dist][x]
	}
	return count
}

func abs2013(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// type DetectSquares struct {
// 	points map[Point2013]int
// 	maxX   int
// 	minX   int
// }

// type Point2013 struct {
// 	x, y int
// }

// func Constructor() DetectSquares {
// 	return DetectSquares{map[Point2013]int{}, 0, 1001}
// }

// var direction = [2]int{-1, 1}

// func (this *DetectSquares) Add(point []int) {
// 	p := Point2013{point[0], point[1]}
// 	if _, ok := this.points[p]; ok {
// 		this.points[p] += 1
// 	} else {
// 		this.points[p] = 1
// 	}

// 	if point[0] > this.maxX {
// 		this.maxX = point[0]
// 	}
// 	if point[0] < this.minX {
// 		this.minX = point[0]
// 	}
// }

// func (this *DetectSquares) Count(point []int) int {
// 	count := 0
// 	x, y := point[0], point[1]
// 	for i := this.minX; i <= this.maxX; i++ {
// 		fmt.Printf("finding for x: %d", i)
// 		if i == x {
// 			continue
// 		}
// 		if p1, ok := this.points[Point2013{i, y}]; ok {
// 			dist := abs2013(i, x)
// 			for _, v := range direction {
// 				if p2, ok := this.points[Point2013{i, y + dist*v}]; ok {
// 					if p3, ok := this.points[Point2013{x, y + dist*v}]; ok {
// 						// fmt.Printf("found points, [%d, %d]=%d, [%d, %d]=%d, [%d, %d]=%d\n",
// 						// 	i, y, p1, i, y+dist*v, p2, x, y+dist*v, p3)
// 						count += p1 * p2 * p3
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// func abs2013(x, y int) int {
// 	if x < y {
// 		return y - x
// 	}
// 	return x - y
// }

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
// @lc code=end
