package Stack_Queue

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 1 {
		return []int{nums[0]}
	}
	// 记录当前队列的最大值和次大值
	res := make([]int, 0, 100000)
	var firstMax, cur = nums[0], 0
	myQueue := NewQueue[int](1000)
	// 入队前k个元素
	for i := 0; i < k; i++ {
		cur = nums[i]
		if cur > firstMax {
			firstMax = cur
		}
		myQueue.Enqueue(cur)
	}

	res = append(res, firstMax) // 将前k个元素中的最大值加入到数组中

	//移动窗口
	for i := k; i < length; i++ {
		//元素出栈
		value, _ := myQueue.Dequeue()
		// 判断当前元素是否为上一个窗口的最大值
		if value == firstMax {
			firstMax = secondMax //当前最大值替换
		}
		// 如果不是则比较新入栈元素与当前最大值
		myQueue.Enqueue(nums[i])
		if nums[i] >= firstMax {
			firstMax = nums[i]
			res = append(res, nums[i])
		} else if nums[i] > secondMax && nums[i] <= firstMax {
			secondMax = nums[i]
		}
		res = append(res, firstMax)
	}
	return res
}
