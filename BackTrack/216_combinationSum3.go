package BackTrack

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var curSum = 0
	var backTracking func(int)
	backTracking = func(start int) {
		// 如果当前长度已满 判断是否等于目标值 如果不等则当前条件不满足
		if len(tmp) == k {
			if curSum == n {
				cp := make([]int, k)
				copy(cp, tmp)
				res = append(res, cp)
				return
			}
			// 表明数组已满 但未达到要求 不必再尝试了 进行回溯剔除
			return
		}
		for i := start; i <= 9-(k-len(tmp))+1; i++ {
			// 如果已有的和加上当前的已经大于n了 则该元素及其后续元素不应该被加入
			if curSum+i > n {
				// 跳出当前循环
				break
			}
			// 否则更新curSum 并且加入待选数组
			curSum += i
			tmp = append(tmp, i)
			backTracking(i + 1)
			// 回溯完后对当前状态进行回溯
			tmp = tmp[:len(tmp)-1] //剔除加入的元素，通过循环换下一个元素
			curSum -= i            // 剔除上一个加入元素的和
		}
	}
	backTracking(1)
	return res
}
