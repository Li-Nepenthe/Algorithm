package Tree

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	allPath("", root, &res)
	return res
}

func allPath(s string, root *TreeNode, res *[]string) {
	if root == nil {
		return
	}
	// 将int转为string
	s = s + strconv.Itoa(root.Val)
	if root.Left != nil || root.Right != nil {
		s = s + "->"
	} else {
		*res = append(*res, s)
	}
	allPath(s, root.Left, res)
	allPath(s, root.Right, res)
	return
}
