package Array

// 思想，最边缘的平方后也是可能最大的，所以不断比较两边的元素挨个填入新数组
func sortedSquares(nums []int) []int {
	var newNums = make([]int, len(nums))
	var i, j = 0, len(nums) - 1
	var k = len(nums) - 1
	for k >= 0 {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			newNums[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			newNums[k] = nums[j] * nums[j]
			k--
			j--
		}
	}
	return newNums
}
