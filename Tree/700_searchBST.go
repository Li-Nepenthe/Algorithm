package Tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		root = searchBST(root.Left, val)
	} else {
		root = searchBST(root.Right, val)
	}
	return root
}
