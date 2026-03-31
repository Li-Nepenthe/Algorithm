package Tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		}
		return root1
	}
	root1.Val = root1.Val + root2.Val
	l, r := mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)
	root1.Left, root1.Right = l, r
	return root1
}
