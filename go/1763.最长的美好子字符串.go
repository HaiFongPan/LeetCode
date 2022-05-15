package leetcode

import "unicode"

/*
 * @lc app=leetcode.cn id=1763 lang=golang
 *
 * [1763] 最长的美好子字符串
 */

// @lc code=start
var pos, lens int

func longestNiceSubstring(s string) string {
	pos = 0
	lens = 0
	search1763(0, len(s)-1, s)
	return s[pos : pos+lens]
}

func search1763(start, end int, s string) {
	if start >= end {
		return
	}
	lower := 0
	upper := 0
	for i := start; i <= end; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			lower |= 1 << (s[i] - 'a')
		} else {
			upper |= 1 << (s[i] - 'A')
		}
	}

	if lower == upper {
		if end-start+1 > lens {
			lens = end - start + 1
			pos = start
		}
		return
	}

	diff := lower & upper
	nPos := start
	for nPos <= end {
		start = nPos
		for nPos <= end && (diff&(1<<(unicode.ToLower(rune(s[nPos]))-'a'))) != 0 {
			nPos++
		}
		search1763(start, nPos-1, s)
		nPos++
	}
}

// @lc code=end
