package Tree

import Stack_Queue "Algorithm/Stack&Queue"

func levelOrderBottom(root *treeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := Stack_Queue.NewQueue[*treeNode](2000)
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		size := queue.Len()
		res = append(res, make([]int, 0, size))
		for i := 0; i < size; i++ {
			//出队就记录在当前层级的数组中
			cur, _ := queue.Dequeue()
			// 直接利用res的长度来记录是当前哪一层 -- 长度为1则记录根节点
			res[len(res)-1] = append(res[len(res)-1], cur.Val)
			if cur.Left != nil {
				queue.Enqueue(cur.Left)
			}
			if cur.Right != nil {
				queue.Enqueue(cur.Right)
			}
		}
	}
	// 将结果反转一下就行
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
