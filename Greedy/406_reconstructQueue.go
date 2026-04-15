package Greedy

import (
	"cmp"
	"slices"
)

// [high,k]表示大于等于high的有k个
// 核心思想就是按照身高降序排序
// 由于场上没有比最高的还高的值 那么他们的k就是应对了当前队列的索引位置
// 由于k只反映前面有没有大于等于它的 矮个子插入在高个子前面并不会有影响
// 基于此 所以本体的贪心思路为优先排高个子的，每次插入的比上一轮的矮就不会影响
func reconstructQueue(people [][]int) [][]int {
	// 按照身高降序排序 然后把矮个子的根据k值插入到前面，相同身高的按照k值升序排序
	// [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
	// [[7,0],[7,1],[6,1],[5,0],[5,2],[4,4]]  这样矮个子对高个子来说不可见
	// 根据新排序的people的k值进行插值排序
	// [7,0]
	// [7,0] [7,1]
	// [7,0] [6,1] [7,1]
	// [5,0] [7,0] [6,1] [7,1]
	// [5,0] [7,0] [5,2] [6,1] [7,1]
	// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	slices.SortFunc(people, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(b[0], a[0]),
			cmp.Compare(a[1], b[1]))
	})
	result := make([][]int, 0, len(people))
	for _, v := range people {
		position := v[1]
		//增加数组长度
		result = append(result, nil)
		//将position后的数组整体向后移动
		// 处理任何一个 [h,k] 当前的result长度一定大于等于k
		copy(result[position+1:], result[position:])
		result[position] = v
	}
	return result
}
