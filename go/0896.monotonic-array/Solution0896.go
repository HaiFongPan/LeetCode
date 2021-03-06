// Package leetcode provides ...
package leetcode

func isMonotonic(A []int) bool {
	flag := 0
	for i := 1; i < len(A); i++ {
		f := A[i-1] - A[i]
		if f == 0 {
			continue
		}
		if f > 0 && flag < 0 {
			return false
		}
		if f < 0 && flag > 0 {
			return false
		}
		flag = f
	}
	return true
}
