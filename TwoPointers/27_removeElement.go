package TwoPointers

func removeElement(nums []int, val int) int {

	// 一个fast指针向前遍历 指向不等于val的元素，
	// slow指向能被替换的位置
	fast, slow := 0, 0
	length := len(nums)
	for fast < length {
		if nums[fast] != val {
			//交换
			nums[slow] = nums[fast]
			fast++
			slow++
		} else {
			fast++
		}
	}
	return slow
}
