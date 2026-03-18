package Array

// 思路：利用快慢指针，fast寻找所有不等于val的元素 slow为fast找到的元素腾出位置

func removeElement(nums []int, val int) int {

	var slow, fast int
	for fast = 0; fast < len(nums); fast++ {
		if nums[slow] == val {
			if nums[fast] != val {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				slow++
			}
		} else {
			slow++
		}
	}
	return slow
}
