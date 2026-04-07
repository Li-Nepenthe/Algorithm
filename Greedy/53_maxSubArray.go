package Greedy

func maxSubArray(nums []int) int {
	// 移动到第一个非零元素下标
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			index = i
		}
	}
	for {

	}
}
