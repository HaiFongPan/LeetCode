package leetcode

/*
 * @lc app=leetcode.cn id=1325 lang=golang
 *
 * [1325] 删除给定值的叶子节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 */
// DFS
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	left := removeLeafNodes(root.Left, target)
	right := removeLeafNodes(root.Right, target)

	root.Left = left
	root.Right = right

	if left == nil && right == nil {
		if root.Val == target {
			return nil
		}
		return root
	}

	return root
}

// @lc code=end
