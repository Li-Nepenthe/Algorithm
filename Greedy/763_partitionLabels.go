package Greedy

func partitionLabels(s string) []int {
	// 1. 预处理：记录每个字符最后出现的索引
	// 使用 [26]int 而不是 make([]int, 26) 在这种固定大小场景下更高效
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	var res []int
	start, end := 0, 0

	// 2. 贪心扫描
	for i := 0; i < len(s); i++ {
		// 更新当前区间必须延伸到的最远边界
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}

		// 当扫描索引追上最远边界时，说明找到了一个最小合法区间
		if i == end {
			res = append(res, end-start+1)
			// 下一个区间的起点是当前终点的下一个位置
			start = i + 1
		}
	}
	return res
}
