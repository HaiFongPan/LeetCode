package leetcode

func findTheLongestSubstring(s string) int {
	max := 0
	xor := 0
	vowel := make([]int, 1<<5)
	for i := 0; i < len(vowel); i++ {
		vowel[i] = -1
	}
	vowel[0] = 0
	for i := 0; i < len(s); i++ {
		if val := isVowel(s[i]); val >= 0 {
			xor ^= val
		}

		if vowel[xor] >= 0 {
			temp := i + 1 - vowel[xor]
			if temp > max {
				max = temp
			}
		} else {
			vowel[xor] = i + 1
		}
	}

	return max
}

func isVowel(c byte) int {
	switch c {
	case 'a':
		return 1 << 0
	case 'e':
		return 1 << 1
	case 'i':
		return 1 << 2
	case 'o':
		return 1 << 3
	case 'u':
		return 1 << 4
	default:
		return -1
	}
}
