package Stack_Queue

// Stack 利用泛型实现栈
type Stack[T any] struct {
	data []T
}

// NewStack 创建一个栈
func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{data: make([]T, 0, cap)}
}

// Push 根据创建时的类型自动替换
func (s *Stack[T]) Push(v T) bool {
	s.data = append(s.data, v)
	return true
}

// Pop 关键改动：返回 (T, bool) 是 Go 处理泛型空值的标准做法
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.data) - 1
	val := s.data[index]

	var zero T
	s.data[index] = zero
	s.data = s.data[:index]

	return val, true
}

func (s *Stack[T]) Clear() {
	// 💡 提示：为了让 GC 回收掉内部元素的引用，
	// 如果 T 是指针类型，最好先遍历置空，或者直接重新 make
	s.data = s.data[:0]
}

func (s *Stack[T]) Cap() int {
	return cap(s.data)
}

func (s *Stack[T]) Shrink() {
	if len(s.data) < cap(s.data)/2 {
		newContainer := make([]T, len(s.data))
		copy(newContainer, s.data)
		s.data = newContainer
	}
}

func (s *Stack[T]) ToSlice() []T {
	res := make([]T, len(s.data))
	copy(res, s.data) // 返回副本，防止外部修改导致栈内部崩溃
	return res
}

func (s *Stack[T]) PeekMany(k int) ([]T, bool) {
	if len(s.data) < k {
		return nil, false
	}
	// 返回栈顶部的 k 个元素（注意顺序是从旧到新还是从新到旧）
	return s.data[len(s.data)-k:], true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
