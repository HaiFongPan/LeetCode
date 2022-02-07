package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=2047 lang=golang
 *
 * [2047] 句子中的有效单词数
 *
 */
// 存在数字不是单词
// 标点不在最后不是单词
// 标点只能有一个
// - 必须在中间

// @lc code=start
func countValidWords(sentence string) int {
	count := 0
	tokens := strings.Split(sentence, " ")
	for _, token := range tokens {
		if len(token) == 0 {
			continue
		}
		tokenLen := len(token)
		hasMinus := false
		for i := 0; i < tokenLen; i++ {
			if token[i] >= '0' && token[i] <= '9' {
				break
			}

			if (token[i] == '!' || token[i] == '.' || token[i] == ',') && (i != tokenLen-1) {
				break
			}
			if token[i] == '-' {
				if i == 0 || i == tokenLen-1 || hasMinus || !isAlpha(token[i+1]) {
					break
				}
				hasMinus = true
			}
			if i == tokenLen-1 {
				count++
			}
		}
	}
	return count
}

func isAlpha(r byte) bool {
	return r >= 'a' && r <= 'z'
}

// @lc code=end
