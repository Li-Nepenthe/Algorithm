package Tree

// 利用后序遍历的回溯性
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点满足任意条件 则直接返回  不再递归
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	// 左右两边都找到信息 则当前节点为公共祖先
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
