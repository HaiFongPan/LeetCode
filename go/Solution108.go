package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	root, index := buildNode(nums)
	tree(root, nums, index)
	return root
}

func tree(node *TreeNode, nums []int, nodeIndex int) {
	if len(nums) == 0 {
		return
	}

	var leftNodeIndex, rightNodeIndex int
	if nodeIndex > 0 {
		left := nums[0:nodeIndex]
		node.Left, leftNodeIndex = buildNode(left)
		tree(node.Left, left, leftNodeIndex)
	}

	if nodeIndex < len(nums)-1 {
		right := nums[nodeIndex+1:]
		node.Right, rightNodeIndex = buildNode(right)
		tree(node.Right, right, rightNodeIndex)
	}
}

func buildNode(nums []int) (*TreeNode, int) {
	index := len(nums) / 2
	node := &TreeNode{Val: nums[index]}
	return node, index
}
