package Stack_Queue

// 算法思想：单调队列（双端队列）
// 核心是利用队列存储最大值的下标来判断是否该出队

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	res := make([]int, 0, length-k+1)
	myQueue := NewQueue[int](k) //队列中存储元素下标
	for i := 0; i < len(nums); i++ {
		// 1. 队尾踢人：只要 nums[i] 比队尾大，就一直弹出队尾
		for !myQueue.IsEmpty() && nums[i] >= nums[myQueue.data[myQueue.Len()-1]] {
			myQueue.data = myQueue.data[:myQueue.Len()-1]
		}
		// 2. 将当前下标 i 推入队尾
		myQueue.Enqueue(i)
		// 3. 队头过期：检查 q[0] 是否超出了窗口范围 [i-k+1, i]
		if i-myQueue.data[0]+1 > k {
			myQueue.Dequeue()
		}
		// 4. 收集答案：如果窗口已经形成，取 nums[q[0]] 放入 res
		if i >= k-1 {
			res = append(res, nums[myQueue.data[0]])
		}
	}
	return res
}
