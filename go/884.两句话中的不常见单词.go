package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	w := make(map[string]int)
	for _, v := range strings.Split(s1, " ") {
		w[string(v)]++
	}
	for _, v := range strings.Split(s2, " ") {
		w[string(v)]++
	}
	var res []string
	for k, v := range w {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// @lc code=end
