package Tree

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode // 追踪中序遍历的前一个节点
	var helper func(*TreeNode) bool
	helper = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !helper(node.Left) {
			return false
		}
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		return helper(node.Right)
	}
	return helper(root)
}
