package leetcode

func subarraySum(nums []int, k int) int {
	sum := make(map[int]int, 0)
	sum[0] = 1
	s, c := 0, 0
	for _, n := range nums {
		s = s + n
		if _, ok := sum[s-k]; ok {
			c += sum[s-k]
		}
		sum[s]++
	}

	return c
}

func subarrayDivide(nums []int, k int) int {
	sum := make(map[int]int, 0)
	sum[0] = 1
	s, c := 0, 0
	for _, n := range nums {
		s = s + n
		div := (s%k + k) % k
		c += sum[div]
		sum[div]++
	}
	return c
}
