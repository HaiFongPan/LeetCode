package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1996 lang=golang
 *
 * [1996] 游戏中弱角色的数量
 */

// @lc code=start
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})
	maxDefense := properties[0][1]
	count := 0
	for _, v := range properties {
		if v[1] < maxDefense {
			count++
		} else {
			maxDefense = v[1]
		}
	}
	return count
}

// @lc code=end
