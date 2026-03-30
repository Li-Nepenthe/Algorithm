package Tree

import (
	Stack_Queue "Algorithm/Stack&Queue"
)

func largestValues(root *treeNode) []int {
	res := make([]int, 0, 100)
	if root == nil {
		return res
	}
	q := Stack_Queue.NewQueue[*treeNode](1000)
	q.Enqueue(root)
	for !q.IsEmpty() {
		size := q.Len()
		var curMax int
		for i := 0; i < size; i++ {
			//取出当前这一层的节点
			cur, _ := q.Dequeue()
			if i == 0 {
				curMax = cur.Val
			}
			if cur.Val > curMax {
				curMax = cur.Val
			}
			if cur.Left != nil {
				q.Enqueue(cur.Left)
			}
			if cur.Right != nil {
				q.Enqueue(cur.Right)
			}
		}
		res = append(res, curMax)
	}
	return res
}
