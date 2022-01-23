package leetcode

import "math/rand"

/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	root *ListNode
}

func Constructor384(head *ListNode) Solution {
	return Solution{root: head}
}

func (this *Solution) GetRandom() (ans int) {
	for p, k := this.root, 1; p != nil; k++ {
		if rand.Intn(k) == 0 {
			ans = p.Val
		}
		p = p.Next
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
