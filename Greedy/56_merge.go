package Greedy

import (
	"cmp"
	"slices"
)

//全部以左边界排序 左边界可以确定最左边的点 而右边界会不断更新 [3,4] [2,6]

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	idx := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[idx][1] >= intervals[i][0] {
			intervals[idx][1] = max(intervals[idx][1], intervals[i][1])
		} else {
			idx++
			intervals[idx] = intervals[i]
		}
	}
	return intervals[:idx+1]
}
