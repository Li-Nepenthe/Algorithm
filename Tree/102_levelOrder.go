package Tree

import Stack_Queue "Algorithm/Stack&Queue"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func levelOrder(root *treeNode) [][]int {
	res := make([][]int, 0) // 这里如果访问res[0]会报错，因为还没定义 为nil
	if root == nil {
		return res
	}
	queue := Stack_Queue.NewQueue[*treeNode](2000)
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		size := queue.Len()
		// 为每一层手动创建一个数组元素
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
	return res
}
