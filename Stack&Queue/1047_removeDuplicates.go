package Stack_Queue

func removeDuplicates(s string) string {
	myStack := NewStack[byte](100)

	for i := range s {
		if !myStack.IsEmpty() {
			if b, _ := myStack.PeekMany(1); b[0] == s[i] {
				myStack.Pop()
				continue
			}
		}
		myStack.Push(s[i])
	}

	b, _ := myStack.PeekMany(myStack.Len())

	return string(b)
}
