package Greedy

import (
	"cmp"
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	count, right := 0, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		//当前的左边界大于等于上一个右边界
		if intervals[i][0] >= right {
			right = intervals[i][1]
			continue
		}
		// 否则就是有交集
		count++
	}
	return count
}
