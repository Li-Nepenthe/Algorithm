package HashTable

func isHappy(n int) bool {

	var slow, fast = n, returnResult(n)

	for fast != slow {
		slow = returnResult(slow)
		fast = returnResult(returnResult(fast))
		if fast == 1 {
			return true
		}
	}
	return slow == 1
}

func returnResult(n int) int {
	sum := 0
	for n != 0 {
		temp := n % 10
		sum += temp * temp
		n /= 10
	}
	return sum
}
