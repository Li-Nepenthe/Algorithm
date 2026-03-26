package Stack_Queue

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{data: make([]int, 0, 100)}
}

func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) Push(v int) bool {
	if len(s.data) == cap(s.data) {
		return false
	}
	s.data = append(s.data, v)
	return true
}

func (s *stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	value := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = 0 //置空
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	InStack  *stack
	OutStack *stack
}

func ConstructorQ() MyQueue {
	return MyQueue{
		InStack:  newStack(),
		OutStack: newStack(),
	}
}

func (queue *MyQueue) Push(x int) {
	queue.InStack.Push(x)
}

func (queue *MyQueue) prepareOut() {
	if queue.OutStack.isEmpty() {
		for !queue.InStack.isEmpty() {
			queue.OutStack.Push(queue.InStack.Pop())
		}
	}
}

func (queue *MyQueue) Pop() int {
	queue.prepareOut()
	return queue.OutStack.Pop()
}

func (queue *MyQueue) Peek() int {
	queue.prepareOut()
	return queue.OutStack.Peek()
}

func (queue *MyQueue) Empty() bool {
	return queue.InStack.isEmpty() && queue.OutStack.isEmpty()
}
