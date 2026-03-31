package Tree

type MNode struct {
	Val   int
	Left  *MNode
	Right *MNode
	Next  *MNode
}

func connect(root *MNode) *MNode {
	if root == nil {
		return root
	}
	//利用切片模拟队列
	q := make([]*MNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			//出队
			cur := q[0]
			q = q[1:]
			// 如果是当前层最后一个 直接next赋为nil
			if i == size-1 {
				cur.Next = nil
			} else {
				// 否则赋值为当前对立的下一个值
				cur.Next = q[0]
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}
