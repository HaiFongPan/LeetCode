// Package main provides ...
package leetcode

func oneEditAway(first string, second string) bool {
	fl, sl := len(first), len(second)
	if fl-sl < -1 || fl-sl > 1 {
		return false
	}
	minl := fl
	if sl < fl {
		minl = sl
	}
	for i := 0; i < minl; i++ {
		if first[i] == second[i] {
			continue
		} else if fl == sl {
			return isSame(first[i+1:], second[i+1:])
		} else if fl < sl {
			return isSame(first[i:], second[i+1:])
		} else {
			return isSame(first[i+1:], second[i:])
		}
	}

	return true
}

func isSame(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	l, r := 0, len(a)-1
	for l <= r {
		if a[l] != b[l] || a[r] != b[r] {
			return false
		}
		l = l + 1
		r = r - 1
	}
	return true
}
