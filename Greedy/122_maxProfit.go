package Greedy

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	diff, sum := 0, 0
	for i := 1; i < len(prices); i++ {
		// 计算收益
		diff = prices[i] - prices[i-1]
		// 如果收益为正则卖出 否则不卖
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}
