package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LH, RH := maxDepth(root.Left), maxDepth(root.Right)
	if LH > RH {
		return 1 + LH
	}
	return 1 + RH
}
