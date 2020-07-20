package leetcode

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if ii, ok := isAlphaAndNumber(s[i]); ok {
			if jj, okk := isAlphaAndNumber(s[j]); okk {
				if ii == jj {
					i++
					j--
				} else {
					return false
				}
			} else {
				j--
			}
		} else {
			i++
		}
	}

	return true
}

func isAlphaAndNumber(b byte) (byte, bool) {
	if b >= '0' && b <= '9' {
		return b, true
	}

	if b >= 'a' && b <= 'z' {
		return b - 32, true
	}

	if b >= 'A' && b <= 'Z' {
		return b, true
	}

	return b, false
}
