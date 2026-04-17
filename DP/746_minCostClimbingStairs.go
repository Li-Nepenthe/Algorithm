package DP

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[len(dp)-1]
}
