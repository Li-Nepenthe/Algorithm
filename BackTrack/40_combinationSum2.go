package BackTrack

import "sort"

// 每个数字在每个组合只能用一次 但是数字可以重复

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int
	var dfs func(start int, currentSum int)
	dfs = func(start int, currentSum int) {
		if currentSum == target {
			res = append(res, append([]int(nil), path...))
			return
		}
		// for循环才是本层的选取节点 也是应该去重的地方
		for i := start; i < len(candidates); i++ {
			if currentSum+candidates[i] > target {
				break
			}
			// 去重逻辑  针对每层相同的元素只调用该元素一层，但是不同层不受影响
			// 如果当前不传入的第一个节点，且该节点和前一个节点相同 则说明该节点重复了
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
