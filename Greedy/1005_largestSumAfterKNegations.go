package Greedy

import (
	"math"
	"slices"
)

func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)
	sum, minAbs := 0, math.MaxInt64
	for _, v := range nums {
		// 找到绝对值最小的数 应对剩余k为奇数但负数全部反转完毕的情况
		if math.Abs(float64(v)) < math.Abs(float64(minAbs)) {
			minAbs = v
		}
		// 先利用k尽可能的反转绝对值最大的负值
		if v < 0 && k > 0 {
			sum -= v
			k--
			continue
		}
		sum += v
	}
	//情况一 负值未完全反转完 但是k已经用完了
	if k == 0 || k%2 == 0 {
		return sum
	}
	if k%2 == 1 {
		// 当前记录的绝对值最小值的值可能为正也可能为负
		if minAbs >= 0 {
			return sum - 2*minAbs
		}
		return sum + 2*minAbs
	}
	return sum
}
