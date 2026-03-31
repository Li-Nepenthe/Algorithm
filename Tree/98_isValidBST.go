package Tree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var f bool
	if root.Left == nil && root.Right == nil {
		f = true
	}
	if root.Left != nil && root.Right != nil {
		f = root.Val > root.Left.Val && root.Val < root.Right.Val
	}
	if root.Left != nil && root.Right == nil {
		f = root.Val > root.Left.Val
	}
	if root.Right != nil && root.Left == nil {
		f = root.Val < root.Right.Val
	}
	return f && isValidBST(root.Left) && isValidBST(root.Right)
}
