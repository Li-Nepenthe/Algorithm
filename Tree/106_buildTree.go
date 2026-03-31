package Tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// 1. 确定根节点值 当前根节点的值永远是后序遍历的最后一个
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 找到根节点在中序的索引
	var index int
	for i, v := range inorder {
		if v == rootVal {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
