package HashTable

func twoSum(nums []int, target int) []int {
	// 1. 预分配容量 = 数组长度，避免在循环中触发 Map 扩容
	m := make(map[int]int, len(nums))

	// 2. 使用 range 获取索引 i 和数值 num
	for i, num := range nums {
		complement := target - num

		// 3. 检查“补数”是否已经在之前的记录中
		if idx, ok := m[complement]; ok {
			// 找到了！idx 是之前存的下标，i 是当前的下标
			return []int{idx, i}
		}

		// 4. 没找到，则记录当前数值的下标，供后续数字匹配
		m[num] = i
	}

	return nil // Go 惯例：未找到返回 nil
}
