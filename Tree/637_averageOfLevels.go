package Tree

import Stack_Queue "Algorithm/Stack&Queue"

func averageOfLevels(root *treeNode) []float64 {
	res := make([]float64, 0)
	queue := Stack_Queue.NewQueue[*treeNode](1000)
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		// 记录当前层级的个数
		size := queue.Len()
		var sum float64
		for i := 0; i < size; i++ {
			cur, _ := queue.Dequeue()
			sum += float64(cur.Val)
			if cur.Left != nil {
				queue.Enqueue(cur.Left)
			}
			if cur.Right != nil {
				queue.Enqueue(cur.Right)
			}
		}
		//存储当前节点的平均值
		res = append(res, sum/float64(size))
	}
	return res
}
