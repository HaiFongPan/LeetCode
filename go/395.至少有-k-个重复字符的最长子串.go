package leetcode

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
var maxLen int

func longestSubstring(s string, k int) int {
	maxLen = 0
	search395(0, len(s)-1, s, k)
	return maxLen
}

func search395(start, end int, s string, k int) {
	if start > end || end-start+1 < k {
		return
	}
	// fmt.Printf("searching %s\n", s[start:end+1])
	letters := make(map[byte]int)
	for i := start; i <= end; i++ {
		letters[s[i]-'a']++
	}
	diff := []int{}
	for k2 := range letters {
		if letters[k2] > 0 && letters[k2] < k {
			diff[k2] = 1
		}
	}

	if len(diff) == 0 {
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
		return
	}

	pos := start
	for pos <= end {
		start = pos
		for pos <= end && diff[s[pos]-'a'] == 0 {
			pos++
		}
		// fmt.Printf("sub searching %d, %d\n", start, pos-1)
		search395(start, pos-1, s, k)
		pos++
	}
}

// @lc code=end
