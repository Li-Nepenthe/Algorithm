package Tree

// 思路：每次都拿中间的节点作为根节点 这也能最大程度保证左右子树的高度均衡

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}
