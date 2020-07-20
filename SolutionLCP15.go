package leetcode

func visitOrder(points [][]int, direction string) []int {
	record := make([]bool, len(points))
	result := make([]int, len(points))
	start := 0

	for i := 0; i < len(record); i++ {
		record[i] = false
	}
	for i := 1; i < len(points); i++ {
		if points[i][0] < points[start][0] {
			start = i
		}
	}

	record[start] = true
	result[0] = start

	for i := 0; i < len(direction); i++ {
		next := -1
		if 'L' == direction[i] {
			for j := 0; j < len(points); j++ {
				if record[j] != true {
					if next == -1 || cross(vec(points[next], points[start]), vec(points[j], points[start])) < 0 {
						next = j
					}
				}
			}
		}

		if 'R' == direction[i] {
			for j := 0; j < len(points); j++ {
				if record[j] != true {
					if next == -1 || cross(vec(points[next], points[start]), vec(points[j], points[start])) > 0 {
						next = j
					}
				}
			}
		}

		record[next] = true
		start = next
		result[i+1] = next
	}

	for i := 0; i < len(points); i++ {
		if record[i] == false {
			result[len(record)-1] = i
			break
		}
	}

	return result
}

func vec(p1 []int, p2 []int) []int {
	v := []int{0, 0}
	v[0] = p1[0] - p2[0]
	v[1] = p1[1] - p2[1]
	return v
}

func cross(v1 []int, v2 []int) int {
	// if > 0, right, < 0 left
	return v1[0]*v2[1] - v1[1]*v2[0]
}
