package Stack_Queue

type Stack struct {
	data []int
}

func newStack() *Stack {
	return &Stack{data: make([]int, 0, 100)}
}

func (s *Stack) len() int {
	return len(s.data)
}

func (s *Stack) Push(v int) bool {
	if len(s.data) == cap(s.data) {
		return false
	}
	s.data = append(s.data, v)
	return true
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	value := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = 0 //置空
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	InStack  *Stack
	OutStack *Stack
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
