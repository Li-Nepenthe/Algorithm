package Stack_Queue

import "strconv"

// strconv.Atoi() ASCII to integer 将一个字符串转为整数
func evalRPN(tokens []string) int {
	myStack := NewStack[int](100)
	var tmp1, tmp2, tmp int
	for _, b := range tokens {
		switch b {
		case "+", "-", "*", "/":
			if b == "+" {
				tmp1, _ = myStack.Pop()
				tmp2, _ = myStack.Pop()
				tmp = tmp1 + tmp2
				myStack.Push(tmp)
			}
			if b == "-" {
				tmp1, _ = myStack.Pop()
				tmp2, _ = myStack.Pop()
				tmp = tmp2 - tmp1
				myStack.Push(tmp)
			}
			if b == "*" {
				tmp1, _ = myStack.Pop()
				tmp2, _ = myStack.Pop()
				tmp = tmp1 * tmp2
				myStack.Push(tmp)
			}
			if b == "/" {
				tmp1, _ = myStack.Pop()
				tmp2, _ = myStack.Pop()
				tmp = tmp2 / tmp1
				myStack.Push(tmp)
			}
		default:
			n, _ := strconv.Atoi(b)
			myStack.Push(n)
		}
	}
	value, _ := myStack.Pop()
	return value
}
