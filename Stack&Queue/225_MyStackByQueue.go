package Stack_Queue

type queue struct {
	data []int
}

func newQueue() *queue {
	return &queue{
		data: make([]int, 0, 100),
	}
}

func (q *queue) Len() int {
	return len(q.data)
}

func (q *queue) Push(x int) {
	if q.Len() == cap(q.data) {
		return
	}
	q.data = append(q.data, x)
}

func (q *queue) Pop() int {
	if q.Len() == 0 {
		return -1
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *queue) Peek() int {
	if q.Len() == 0 {
		return -1
	}
	return q.data[0]
}

func (q *queue) Empty() bool {
	return q.Len() == 0
}

type MyStack struct {
	InQueue  *queue
	OutQueue *queue
}

func Constructor() MyStack {
	return MyStack{
		InQueue:  newQueue(),
		OutQueue: newQueue(),
	}
}

func (stack *MyStack) Push(x int) {
	stack.InQueue.Push(x)
}

func (stack *MyStack) Pop() int {
	if stack.InQueue.Len() == 1 {
		return stack.InQueue.Pop()
	}
	for stack.InQueue.Len() > 1 {
		stack.OutQueue.Push(stack.InQueue.Pop())
	}
	value := stack.InQueue.Pop()
	stack.InQueue, stack.OutQueue = stack.OutQueue, stack.InQueue
	return value
}

func (stack *MyStack) Top() int {
	if stack.InQueue.Len() == 1 {
		return stack.InQueue.Peek()
	}
	for stack.InQueue.Len() > 1 {
		stack.OutQueue.Push(stack.InQueue.Pop())
	}
	value := stack.InQueue.Pop()
	stack.OutQueue.Push(value)
	stack.InQueue, stack.OutQueue = stack.OutQueue, stack.InQueue
	return value
}

func (stack *MyStack) Empty() bool {
	return stack.InQueue.Empty()
}
