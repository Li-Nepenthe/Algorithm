package Tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Val, root.Right.Val = root.Right.Val, root.Left.Val
	}
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
