package BackTrack

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var backTrack func(int)
	backTrack = func(start int) {
		if len(tmp) == k {
			cp := make([]int, k)
			copy(cp, tmp) // 这一步必不可少！
			res = append(res, cp)
			return
		}
		// n-(k-len(tmp))+1 每个元素都能遍历到
		for i := start; i <= n-(k-len(tmp))+1; i++ {
			tmp = append(tmp, i)
			backTrack(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backTrack(1)
	return res
}
