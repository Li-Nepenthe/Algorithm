package Greedy

import (
	"slices"
)

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var count, i, j = 0, len(s) - 1, len(g) - 1
	// 用最大的饼干满足胃口最大的
	for i >= 0 && j >= 0 {
		if s[i] >= g[j] {
			i--
			j--
			count++
		} else {
			j--
		}
	}
	return count
}
