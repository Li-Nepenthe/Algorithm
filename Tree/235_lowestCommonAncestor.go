package Tree

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right // 向右移
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left // 向左移
		} else {
			return root // 找到分叉点或匹配节点
		}
	}

	return root
}
