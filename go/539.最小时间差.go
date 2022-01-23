package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	sort.Strings(timePoints)
	prev := toMinute(timePoints[0])
	tail := toMinute(timePoints[len(timePoints)-1])
	min := 1440 - tail + prev
	for i := 1; i < len(timePoints); i++ {
		curr := toMinute(timePoints[i])
		if curr-prev < min {
			min = curr - prev
		}
		prev = curr
	}
	return min
}

func toMinute(s string) int {
	return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + (int(s[3]-'0')*10 + int(s[4]-'0'))
}

// @lc code=end
