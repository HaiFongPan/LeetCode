package leetcode

import "strings"

func longestSubstring(s string, k int) (ans int) {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	var split byte
	for i := 0; i < 26; i++ {
		c := cnt[i]
		if c > 0 && c < k {
			split = 'a' + byte(c)
		}
	}

	if split == 0 {
		return len(s)
	}

	for _, v := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(v, k))
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
