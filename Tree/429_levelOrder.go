package Tree

import Stack_Queue "Algorithm/Stack&Queue"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderForN(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := Stack_Queue.NewQueue[*Node](1000)
	q.Enqueue(root)
	for !q.IsEmpty() {
		size := q.Len()
		res = append(res, make([]int, 0, size))
		for i := 0; i < size; i++ {
			//取出当前这一层的节点
			cur, _ := q.Dequeue()
			res[len(res)-1] = append(res[len(res)-1], cur.Val)
			//将当前节点下所有子节点全部入队
			for _, v := range cur.Children {
				q.Enqueue(v)
			}
		}
	}
	return res
}
