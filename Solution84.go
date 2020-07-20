package leetcode

func largestRectangleArea(heights []int) int {
	result := 0

	for i := 0; i < len(heights); i++ {
		area := heights[i]
		maxHeight := heights[i]

		for j := i + 1; j < len(heights) && heights[j] > 0; j++ {
			maxHeight = Min(maxHeight, heights[j])
			width := j - i + 1
			area = Maxx(area, width*maxHeight)
		}

		result = Maxx(area, result)
	}
	return result
}

func Maxx(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
