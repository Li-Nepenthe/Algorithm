package Tree

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, f := getHigh(root)
	return f
}

func getHigh(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	LH, LF := getHigh(root.Left)
	RH, RF := getHigh(root.Right)
	flag := LF && RF
	if math.Abs(float64(LH)-float64(RH)) <= 1 {
		return 1 + max(LH, RH), flag
	}
	return 1 + max(LH, RH), false
}
