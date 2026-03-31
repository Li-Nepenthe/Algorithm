package Tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	curMax, index := 0, 0
	for i, num := range nums {
		if num > curMax {
			curMax = num
			index = i
		}
	}
	root := &TreeNode{Val: curMax}
	l := constructMaximumBinaryTree(nums[:index])
	r := constructMaximumBinaryTree(nums[index+1:])
	root.Left = l
	root.Right = r
	return root
}
