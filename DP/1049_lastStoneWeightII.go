package DP

import "fmt"

// 本题的思想就是凑出两个阵营 两边重量相差最小即可
// 每个越接近一半 最后碰撞出的就是最小
// 不要每个石头每个石头看 当作一个整体看待

func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, v := range stones {
		for i := target; i >= v; i-- {
			dp[i] = dp[i] || dp[i-v]
		}
	}
	for i := target; i >= 0; i-- {
		if dp[i] {
			fmt.Println(i)
			return sum - i*2
		}
	}
	return -1
}
