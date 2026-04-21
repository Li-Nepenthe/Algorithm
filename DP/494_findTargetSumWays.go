package DP

// 本质上就是把数组中的缘分分为正数P和负数N两个集合
// 于是有P-N = Target P+N = Sum(nums)
// 二者相加得 2*P = Sum(nums) + Target 公式一
// --> P = (Sum(nums) + Target)/2
// --> N = (Sum(nums) - Target)/2
// 同时如果target大于Sum(nums)则凑不出
// 注意公式一等式左边为偶数 所以如果Sum(nums) + Target 不为偶数 则也不可能凑出来

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if (sum-target < 0) || (sum+target)%2 != 0 {
		return 0
	}

	newTarget := (sum - target) / 2
	dp := make([]int, newTarget+1)
	dp[0] = 1
	for _, num := range nums {
		for i := newTarget; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[newTarget]
}
