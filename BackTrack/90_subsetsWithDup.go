package BackTrack

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// 先对元素进行排序
	sort.Ints(nums)
	res := make([][]int, 0)
	path := make([]int, 0)
	var backTracking func(int)
	backTracking = func(start int) {
		cp := make([]int, len(path))
		copy(cp, path)
		res = append(res, cp)
		for i := start; i < len(nums); i++ {
			// 同一层逻辑去重
			if i > start && nums[i] == nums[i-1] {
				// 跳过当前的i
				continue
			}
			// 否则加入path并递归
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}
