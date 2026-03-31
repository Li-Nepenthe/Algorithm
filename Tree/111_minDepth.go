package Tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LH, RH := minDepth(root.Left), minDepth(root.Right)
	if LH == 0 && RH == 0 {
		return 1
	}
	if LH == 0 || RH == 0 {
		return 1 + max(LH, RH)
	}
	return 1 + min(LH, RH)
}
