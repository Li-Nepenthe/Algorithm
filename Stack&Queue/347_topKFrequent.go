package Stack_Queue

// Item 存储元素及其频率
type Item struct {
	value    int // 节点的值
	priority int // 频率
}

// PriorityQueue 实现 heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// 注意：这里决定了是“大顶堆”还是“小顶堆”
	// 如果是 pq[i].priority < pq[j].priority，则是小顶堆（根部是最小元素）
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push 和 Pop 使用指针接收者，因为它们会修改切片的长度
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
