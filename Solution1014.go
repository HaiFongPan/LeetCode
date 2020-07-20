package leetcode

func maxScoreSightseeingPair(A []int) int {
	max := A[0]
	result := 0
	for j := 1; j < len(A); j++ {
		v := A[j] - j
		if v+max > result {
			result = v + max
		}
		if A[j]+j > max {
			max = A[j] + j
		}

	}

	return result
}
