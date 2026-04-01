package Tree

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	var prev *TreeNode
	var maxCount, curCount = 1, 0
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if prev != nil && node.Val == prev.Val {
			curCount++
		} else {
			curCount = 1
		}
		if curCount > maxCount {
			maxCount = curCount
			res = res[:0]
			res = append(res, node.Val)
		} else if curCount == maxCount {
			res = append(res, node.Val)
		}
		prev = node
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}
