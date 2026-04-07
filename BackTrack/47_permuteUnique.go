package BackTrack

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 1. 排序是去重的基础
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))
	length := len(nums)
	used := [8]bool{} // 假设长度不超过 8

	var backTracking func()
	backTracking = func() {
		if len(path) == length {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		for i := 0; i < length; i++ {
			// 场景 A: 树枝去重 (索引 i 已经被用在当前 path 里了)
			if used[i] {
				continue
			}
			// 场景 B: 树层去重 (重点！)
			// 如果当前数等于前一个数，且前一个数【没被使用】
			// 说明前一个数刚刚被回溯掉，现在是同一层在尝试第二个相同的数
			// 因为i转移到这个数说明已经回溯了 一定标记为未使用
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backTracking()

			// 回溯
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTracking()
	return res
}
