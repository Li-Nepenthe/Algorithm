package DP

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	tNumber := n / 3
	left := n - 3*tNumber
	if left == 1 {
		return pow(tNumber-1) * 4
	}
	if left == 0 {
		return pow(tNumber)
	}
	return pow(tNumber) * left
}

func pow(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum = sum * 3
	}
	return sum
}

func integerBreakByDP(n int) int {

}
