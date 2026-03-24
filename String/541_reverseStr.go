package String

func reverseStr(s string, k int) string {
	nums := []byte(s)
	// i每次移动2k个单位
	for i := 0; i < len(s); i += 2 * k {
		end := min(i+k, len(s)) //判断i+k是否超过数组长度 如果超过则表明剩余数组不足k则反转剩余所有的
		swap(nums[i:end])       //否则必定反转前k个
	}
	return string(nums)
}

func swap(nums []byte) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
