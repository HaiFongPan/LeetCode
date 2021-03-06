package leetcode

func validPalindrome(str string) bool {
	is, left, right := isPalindromeSub(str, 0, len(str)-1)
	if !is {
		is, _, _ = isPalindromeSub(str, left+1, right)
		if !is {
			is, _, _ = isPalindromeSub(str, left, right-1)
		}
	}
	return is
}

func isPalindromeSub(str string, left, right int) (bool, int, int) {
	for str[left] == str[right] && left < right {
		left++
		right--
	}
	return left >= right, left, right
}
