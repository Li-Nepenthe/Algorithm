package Tree

// 思路：每次都拿中间的节点作为根节点 这也能最大程度保证左右子树的高度均衡

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	var r *TreeNode
	var helper func([]int, *TreeNode) *TreeNode
	helper = func(num []int, root *TreeNode) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		if len(nums) == 1 {
			return &TreeNode{Val: nums[0], Left: nil, Right: nil}
		}
		if len(nums) == length {
			r = &TreeNode{Val: nums[length/2], Left: nil, Right: nil}
			root = r
		}
		root.Left = helper(nums[:len(num)/2], root.Left)
		root.Right = helper(nums[len(num)/2:], root.Right)
		return root
	}
	return helper(nums, nil)
}
