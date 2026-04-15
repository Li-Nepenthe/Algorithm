package Greedy

import "math"

func maxSubArray(nums []int) int {
	//由于可能出现数组全是负数的情况，所以不能直接跳过负数

	//用一个变量记录当前的和，一个变量记录最大的和
	curSum, maxSum := 0, math.MinInt64
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		if maxSum < curSum {
			maxSum = curSum
		}
		// 如果当前的累计和已经小于等于0了 之前的不在再尝试
		if curSum <= 0 {
			curSum = 0
		}
	}
	return maxSum
}
