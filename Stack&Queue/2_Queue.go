package Stack_Queue

// Queue 利用泛型实现队列
type Queue[T any] struct {
	data []T
}

// NewQueue 创建一个队列
func NewQueue[T any](cap int) *Queue[T] {
	return &Queue[T]{data: make([]T, 0, cap)}
}

// Enqueue 入队：在队尾添加元素
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue 出队：从队头移除并返回元素
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	val := q.data[0]

	// 🧠 工业级细节：手动置空队头，防止内存泄露 (GC 优化)
	q.data[0] = zero
	q.data = q.data[1:]

	return val, true
}

// Peek 查看队头元素但不移除
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.data[0], true
}

// Len 返回队列长度
func (q *Queue[T]) Len() int {
	return len(q.data)
}

// IsEmpty 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Clear 清空队列
func (q *Queue[T]) Clear() {
	var zero T
	// 这里的遍历清空对于存储指针的队列至关重要
	for i := range q.data {
		q.data[i] = zero
	}
	q.data = q.data[:0]
}

// ToSlice 返回队列的切片副本
func (q *Queue[T]) ToSlice() []T {
	res := make([]T, len(q.data))
	copy(res, q.data)
	return res
}
