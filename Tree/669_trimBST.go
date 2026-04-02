package Tree

// 对于每一个节点只可能有三种情况
// 1.该节点的值在范围之前，不管，递归处理左右子节点
// 2.该节点的值小于最小范围，于是该节点及左孩子全部over
// 3.该节点的值大于最大范围，于是该节点及右孩子全部over

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 检验情况一
	if root.Val >= low && root.Val <= high {
		return &TreeNode{
			Val:   root.Val,
			Left:  trimBST(root.Left, low, high),
			Right: trimBST(root.Right, low, high),
		}
	}
	// 检验情况二
	if root.Val < low {
		root = root.Right
		return trimBST(root, low, high)
	}
	// 检验情况三
	if root.Val > high {
		root = root.Left
		return trimBST(root, low, high)
	}
	return nil
}
