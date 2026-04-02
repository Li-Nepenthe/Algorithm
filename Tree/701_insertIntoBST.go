package Tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	r := root
	var prev *TreeNode
	for root != nil {
		prev = root
		if root.Val < val {
			root = root.Right
		} else if root.Val > val {
			root = root.Left
		}
	}
	if prev.Val > val {
		prev.Left = &TreeNode{
			Val: val,
		}
	} else {
		prev.Right = &TreeNode{
			Val: val,
		}
	}

	return r
}
