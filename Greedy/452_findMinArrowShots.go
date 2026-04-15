package Greedy

import (
	"cmp"
	"slices"
)

func findMinArrowShots(points [][]int) int {

	//按照左边界排序 points = [[10,16],[2,8],[1,6],[7,12]]
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	// [1,6] [2,8] [7,12] [10,16]
	// 求两个之间是否有交集
	cross, arrow := []int{points[0][0], points[0][1]}, 1
	for _, v := range points {
		// 如果二者之前没有交集
		if cross[1] < v[0] {
			cross[0], cross[1], arrow = v[0], v[1], arrow+1
			continue
		}
		//更新左右边界
		cross[0], cross[1] = max(cross[0], v[0]), min(cross[1], v[1])
	}
	return arrow
}

func findMinArrowShotsImprove(points [][]int) int {
	//可以按照右边界来排序
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	// 第一支箭射在第一个气球的最右端
	arrow := 1
	lastPos := points[0][1]
	for i := 0; i < len(points); i++ {
		// 如果当前的气球左边界都大于上一个的右边界 则绝对射不到  需要一根新的arrow
		if points[i][0] > lastPos {
			arrow++
			// 更新最大能射到的位置
			lastPos = points[i][1]
		}
	}
	return arrow
}
