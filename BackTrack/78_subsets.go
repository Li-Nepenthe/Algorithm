package BackTrack

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	// 先声明变量类型 再给变量赋值
	var backTracking func(int)
	backTracking = func(start int) {
		cp := make([]int, len(path))
		copy(cp, path)
		res = append(res, cp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}
