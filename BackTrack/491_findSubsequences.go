package BackTrack

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backTracking func(int)
	backTracking = func(start int) {
		if len(path) >= 2 {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
		}
		used := [201]bool{}
		for i := start; i < len(nums); i++ {
			// 同一层要去重 由于是无序的 所以引入used数组
			if used[nums[i]+100] {
				// 跳过当前的i
				continue
			}
			// 还要判断大小
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			// 否则加入path并递归
			path = append(path, nums[i])
			used[nums[i]+100] = true
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}
