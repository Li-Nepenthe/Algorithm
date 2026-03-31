package Tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := make([]*TreeNode, 0, 50)
	q = append(q, root.Left)
	q = append(q, root.Right)
	for len(q) != 0 {
		l, r := q[0], q[1]
		q = q[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, l.Left)
		q = append(q, r.Right)
		q = append(q, l.Right)
		q = append(q, r.Left)
	}
	return true
}
