package Stack_Queue

import (
	"fmt"
	"sort"
)

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
	var data []int
	//移动窗口
	for i := k; i < length; i++ {
		myQueue.Dequeue()        //出
		myQueue.Enqueue(nums[i]) //入
		data = myQueue.data[:myQueue.Len()]
		sort.Ints(data) //找到最大值
		fmt.Println(data)
		fmt.Println(myQueue.data)
		res = append(res, data[myQueue.Len()-1])
	}
	return res
}
