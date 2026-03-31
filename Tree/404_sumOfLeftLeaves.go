package Tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 判断当前节点是否是叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
