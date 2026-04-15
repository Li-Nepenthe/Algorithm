package Greedy

// 维护两个边界，当前的边界，以及边界内能跳的最大的点
// 每次只在边界内找到最大的点跳跃 而不是只要能跳的更远就跳跃
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	step := 0
	maxReach := 0 //maxReach记录的是能跳的最远的绝对坐标而不是相对距离
	end := 0      // 当前这一跳能到的最远边界
	// 注意：我们不需要遍历最后一个元素，因为到达它之前我们已经知道结果了
	for i := 0; i < len(nums)-1; i++ {
		// 1. 贪心地更新全局最远能到的地方
		maxReach = max(maxReach, i+nums[i])
		// 2. 如果走到了当前这一跳的尽头
		if i == end { //当i = 0时 当前操作就是第一步必须跳一次，每次都在结算上一个的跳跃而不是跳到下一个点
			step++         // 必须再跳一次
			end = maxReach // 下一跳的边界就是我们刚才观察到的最远点
		}
	}
	return step
}
