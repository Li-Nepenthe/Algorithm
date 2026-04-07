package Greedy

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var diff, prevDiff int
	var length int
	for i := 1; i < len(nums); i++ {
		diff = nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 {
			length++
			prevDiff = diff
		}
		if diff < 0 && prevDiff >= 0 {
			length++
			prevDiff = diff
		}
	}
	return length + 1
}
