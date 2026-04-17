package Greedy

func monotoneIncreasingDigits(n int) int {
	//要想数最大 最高位必须最大
	number := make([]int, 0)
	length, last, flag := 0, 0, true
	for n > 0 {
		if last > n%10 {
			flag = false
		}
		last = n % 10
		number = append(number, last)
		n = n / 10
		length++
	}
}

func getMin(a int, b int) int {

