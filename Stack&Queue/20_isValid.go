package Stack_Queue

func isValid(s string) bool {
	myStack := NewStack[byte](100)

	for i := range s {
		switch s[i] {
		case '[':
			myStack.Push(s[i])
		case '(':
			myStack.Push(s[i])
		case '{':
			myStack.Push(s[i])
		case ']':
			if b, _ := myStack.Pop(); b != '[' {
				return false
			}
		case ')':
			if b, _ := myStack.Pop(); b != '(' {
				return false
			}
		case '}':
			if b, _ := myStack.Pop(); b != '{' {
				return false
			}
		}
	}

	return myStack.IsEmpty()
}
