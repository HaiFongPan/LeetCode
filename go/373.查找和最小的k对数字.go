package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start
func kSmallestPairsOn2(nums1 []int, nums2 []int, k int) [][]int {
	r := [][]int{}
	for i := 0; i < k && i < len(nums1); i++ {
		for j := 0; j < k && j < len(nums2); j++ {
			r = append(r, []int{nums1[i], nums2[j]})
		}
	}

	if len(r) <= k {
		return r
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i][0]+r[i][1] < r[j][0]+r[j][1]
	})

	return r[0:k]
}

// @lc code=end
