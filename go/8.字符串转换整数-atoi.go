package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	const (
		floor = -2147483648
		ceil  = 2147483647
		blank = ' '
		minus = '-'
		plus  = '+'
	)

	start := false
	candi := []rune{}
	positive := 1
	for _, v := range s {
		if !start && (v == blank || !(v >= '0' && v <= '9')) {
			if v == minus {
				positive = -1
				start = true
			} else if v == plus {
				start = true
			} else if v != blank {
				break
			}
			continue
		}

		if v >= '0' && v <= '9' {
			start = true
			if v == '0' && len(candi) == 0 {
				continue
			}
			candi = append(candi, v)
		} else {
			break
		}
	}
	n := len(candi)
	if n == 0 {
		return 0
	}
	if n > 10 {
		if positive == 1 {
			return ceil
		} else {
			return floor
		}
	}
	r := 0
	for i := 0; i < len(candi); i++ {
		r = r + (int(candi[i]-'0') * int(math.Pow10(n-i-1)) * positive)
	}

	if r < floor {
		return floor
	}
	if r > ceil {
		return ceil
	}

	return r
}

// @lc code=end
