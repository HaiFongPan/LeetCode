package leetcode

func maxNumberOfBalloons(text string) int {
	letters := map[int]int{}
	// a, b, l,o,n
	set := map[int]bool{0: true, 1: true, 11: true, 14: true, 13: true}
	for _, v := range text {
		temp := int(v - 'a')
		if set[temp] {
			letters[temp]++
		}
	}
	min := 10000
	for k, v := range letters {
		cnt := v
		if k == 11 || k == 14 {
			cnt = cnt / 2
		}
		if cnt < min {
			min = cnt
		}
	}
	return min
}
