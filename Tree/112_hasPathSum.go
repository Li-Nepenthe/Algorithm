package Tree

// 这里通过target - 当前节点值获取最终叶节点应该取得的值
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 如果当前节点为叶子节点 判断是否等于最终目标
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	// 递归左右子树，只要有一条路径满足即可（利用 || 的短路特性）
	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}
