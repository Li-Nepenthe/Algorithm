package Stack_Queue

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0, 100),
	}
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Push(x int) {
	if q.Len() == cap(q.data) {
		return
	}
	q.data = append(q.data, x)
}

func (q *Queue) Pop() int {
	if q.Len() == 0 {
		return -1
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *Queue) Peek() int {
	if q.Len() == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

type MyStack struct {
	InQueue  *Queue
	OutQueue *Queue
}

func Constructor() MyStack {
	return MyStack{
		InQueue:  NewQueue(),
		OutQueue: NewQueue(),
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
