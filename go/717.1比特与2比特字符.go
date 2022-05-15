package leetcode

func isOneBitCharacter(bits []int) bool {
	// 1,1,0
	i := 0
	ans := false
	for i < len(bits) {
		if bits[i] == 1 {
			ans = false
			i += 2
		} else {
			ans = true
			i++
		}
	}
	return ans
}
