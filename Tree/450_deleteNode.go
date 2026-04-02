package Tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 如果当前节点为空 则直接返回nil
	if root == nil {
		return nil
	}
	// 如果被删除的值小于当前节点值，则必处于当前节点的左孩子
	if key < root.Val {
		// 当前节点的左孩子更新为删除完指定节点后的装填
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		// 当前节点的右孩子更新为删除完指定节点后的装填
		root.Right = deleteNode(root.Right, key)
	} else { // 当前节点为被删除节点
		// 如果为叶子节点，则直接置为空并返回
		if root.Left == nil && root.Right == nil {
			root = nil
			return root
		}
		// 如果当前节点的左孩子为空，则直接返回当前右孩子作为新的根节点
		if root.Left == nil {
			return root.Right
		}
		// 如果当前节点右孩子为空，则直接返回当前左孩子为新的根节点
		if root.Right == nil {
			return root.Left
		}
		// 如果当前节点既有左孩子又有右孩子，要么找到左边孩子节点的最大值即直接前驱，要么找到右边孩子的最小值即直接后继
		// 直接前驱和直接后继不可能再有两个孩子的情况，要么为叶子节点要么为单孩子节点
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)

	}
	return root
}
