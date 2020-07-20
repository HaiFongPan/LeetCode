package leetcode

var s int

func convertBST(root *TreeNode) *TreeNode {
	s = 0
	gcd(root)
	return root
}

func gcd(node *TreeNode) {
	if node != nil {
		gcd(node.Right)
		s += node.Val
		node.Val = s
		gcd(node.Left)
	}
}
