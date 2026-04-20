package DP

// 计算N时考虑1到N每一个都为根节点的情况
// 假如我们选取第j个元素  1 <= j <= N
// 以j为根元素的种类情况就是 所有比j小的元素作为j的左子树的种类X所有比j大的元素作为右子树的种类
// 即DP[j] = DP[j-1]*DP[N-j]

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		sum := 0
		// 从1开始依次为根节点的情况
		for j := 1; j <= i; j++ {
			sum += dp[j-1] * dp[i-j]
		}
		dp[i] = sum
	}
	return dp[n]
}
