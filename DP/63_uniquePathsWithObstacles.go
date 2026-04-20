package DP

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 1.确定DP以及下标的含义

	//第1列和第1行 遇到的非障碍全部置为1 表示只有-1一种走法 如果遇到障碍，则后面的全部置为-1 表示不可达
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	//2.确定递推公式
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			//当前节点面临可能的情况

			// 如果当前节点本身就是障碍，直接置为-1
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = -1
				continue
			}
			//左边和上边都能到达当前位置
			if dp[i-1][j] != -1 && dp[i][j-1] != -1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
				continue
			}
			//左和上两边都不能到达当前位置
			if dp[i-1][j] == -1 && dp[i][j-1] == -1 {
				dp[i][j] = -1
				continue
			}
			// 左右两边有一个能到达
			if dp[i][j-1] != -1 {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	if dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == -1 {
		return 0
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func uniquePathsWithObstaclesImprove(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([]int, n)
	dp[0] = 1 //dp[0]表示不可达
	// 按行优先遍历
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前节点是障碍 则表示此路不通 置为0
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			//  dp[j] += dp[j-1] --> dp[j] = dp[j]+ dp[j-1]
			// 其中左边的dp[j]表示当前行的状态
			// 右边的dp[j]表示上一行的dp状态
			// dp[j-1]表示当前行左边的状态

			// 通过0值 覆盖了路径不同和可以走通的情况 路径不通则提供0贡献
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
