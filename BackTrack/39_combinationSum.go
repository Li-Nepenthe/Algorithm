package BackTrack

import (
	"sort"
)

func combinationSum1(candidates []int, target int) [][]int {
	// 先让数组中的元素有序
	sort.Ints(candidates)
	// 题目规定数组长度在1到30之间
	if candidates[0] > target {
		return [][]int{}
	}
	if candidates[0] == target {
		return [][]int{
			0: {candidates[0]},
		}
	}
	res := make([][]int, 0, 150)
	tmp := make([]int, 0)
	var sum, prev int
	var backTracking = func() {}
	backTracking = func() {
		// 如果当前的和满足要求则返回 并且保存在res结果数组中
		if sum == target {
			cp := make([]int, len(tmp))
			copy(cp, tmp)
			res = append(res, cp)
			return
		}
		for i := 0; i < len(candidates); i++ {
			// 在累加之前统计一下，如果当前累加后的值已经大于target 由于有序排序且无重复，
			//  后续值更不可能 所以要跳出循环
			if sum+candidates[i] > target {
				return
			}
			// 如果当前节点的值小于上一层的节点 说明会重复 直接跳过
			if candidates[i] < prev {
				continue
			}
			sum += candidates[i]
			tmp = append(tmp, candidates[i])
			last := prev         //保存上一层prev 用于回溯
			prev = candidates[i] // 更新prev
			backTracking()
			// 回溯
			sum -= candidates[i]
			tmp = tmp[:len(tmp)-1]
			prev = last
		}
	}
	backTracking()
	return res
}

func combinationSum1Improve(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int
	var dfs func(start int, currentSum int)
	dfs = func(start int, currentSum int) {
		if currentSum == target {
			res = append(res, append([]int(nil), path...)) // 优雅的 Copy 方式
			return
		}

		for i := start; i < len(candidates); i++ {
			if currentSum+candidates[i] > target {
				// 跳出循环
				break
			}
			path = append(path, candidates[i])
			// 由于可以无限制的调用本身 所以这里不传递i+1 而是从自己开始
			dfs(i, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
