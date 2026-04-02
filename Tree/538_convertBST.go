package Tree

// 思路，先计算所有节点的值的和allSum 然后中序遍历分别用allSum减去上一个节点的值即位当前的累加和
// 进阶，如果采用逆中序遍历，返回的就是从大到小的值

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 1. 先去最右边（最大的地方）
		dfs(node.Right)

		// 2. 更新全局累加和并修改当前节点
		sum += node.Val
		node.Val = sum

		// 3. 再去左边（较小的地方）
		dfs(node.Left)
	}
	dfs(root)
	return root
}
