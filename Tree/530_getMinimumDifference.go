package Tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	// 初始化最小差值为最大整数
	minDiff := math.MaxInt
	// prev 指向遍历过程中的前一个节点
	var prev *TreeNode
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 1. 递归左子树
		helper(node.Left)
		// 2. 处理当前节点
		if prev != nil {
			// 计算当前节点与前一个节点的差值，并更新最小值
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		// 更新 prev 为当前节点，供下一次比较使用
		prev = node
		// 3. 递归右子树
		helper(node.Right)
	}
	helper(root)
	return minDiff
}
