package Array

// 利用滑动窗口来做

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var length = len(nums) + 1 //作为不可达数
	var start, currentSum int
	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]
		// 如果累计和大于target 则尝试不断缩小左边界
		for currentSum >= target {
			var currentLength = end - start + 1
			// 如果长度小于当前长度 才缩小length
			if currentLength < length {
				length = currentLength //更新最小长度
			}
			currentSum -= nums[start] //尝试减小范围
			start++
		}
	}

	//如果未进循环 则说明不满足 length为不可达数
	if length == len(nums)+1 {
		return 0
	}

	return length
}
