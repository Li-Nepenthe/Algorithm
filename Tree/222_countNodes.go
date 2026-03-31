package Tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lCount := countNodes(root.Left)
	rCount := countNodes(root.Right)
	return 1 + lCount + rCount
}
