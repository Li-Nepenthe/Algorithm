package BackTrack

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	length := len(nums)
	used := [21]bool{}
	var backTracking func([]int)
	backTracking = func(ints []int) {
		if len(path) == length {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}
		for i := 0; i < length; i++ {
			// 当前数组用过 直接跳过
			if used[nums[i]+10] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]+10] = true
			backTracking(nums)
			path = path[:len(path)-1]
			used[nums[i]+10] = false
		}

	}

	backTracking(nums)

	return res
}
