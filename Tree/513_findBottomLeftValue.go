package Tree

func findBottomLeftValue(root *TreeNode) int {
	q := make([]*TreeNode, 0, 100)
	q = append(q, root)
	var res int
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if i == 0 {
				res = cur.Val
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return res
}
