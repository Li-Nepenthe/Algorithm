package BackTrack

// 本题的核心在于针对每一层，得回溯不同层次，比如选取前一个，前两个，前三个等等
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	//检验当前的字符串数组的字符串元素是否满足回文串定义
	var isValid func(s string) bool
	isValid = func(path string) bool {
		length := len(path)
		if length == 0 {
			return false
		}
		for i := 0; i < length/2; i++ {
			if path[i] != path[length-i-1] {
				return false
			}
		}
		return true
	}
	var backTracking func(string)
	backTracking = func(s string) {
		// 如果传进来的数组为0 则表示走到最后一个了 将当前所有结果保存
		if len(s) == 0 {
			cp := make([]string, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}
		for i := 0; i < len(s); i++ {
			// 尝试不同长度的切割
			if isValid(s[:i+1]) {
				//在路径中增加
				path = append(path, s[:i+1])
				// 尝试递归
				// 如果当前长度的切割为回文 则把剩下的数组传递进去递归
				backTracking(s[i+1:])
				// 递归后回复原样
				// 对路径函数进行回溯
				path = path[:len(path)-1]
			}
			continue
		}
	}
	backTracking(s)
	return res
}

func partitionImprove(s string) [][]string {
	n := len(s)
	// 假设 n 是字符串长度，s 是原始字符串
	isPali := make([][]bool, n)
	for i := range isPali {
		isPali[i] = make([]bool, n)
	}

	// 从下往上遍历 i，从左往右遍历 j
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// 核心判断逻辑
			if s[i] == s[j] {
				// 情况 1：子串长度为 1 (i == j) 或 2 (j - i == 1) 或 3 (j - i == 2)
				// 只要首尾相等，它们一定是回文
				if j-i <= 2 {
					isPali[i][j] = true
				} else {
					// 情况 2：长度大于 3，取决于内部子串 [i+1, j-1] 是否为回文
					isPali[i][j] = isPali[i+1][j-1]
				}
			}
		}
	}
	res := make([][]string, 0)
	path := make([]string, 0)
	var backTracking func(int)
	backTracking = func(start int) {
		// 终止条件：start 走到了原始字符串的末尾
		if start == n {
			cp := make([]string, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		// i 是当前的结束位置（绝对坐标）
		for i := start; i < n; i++ {
			// 🎯 直接查表：从 start 到 i 是否为回文
			if isPali[start][i] {
				// 记录路径：利用原始字符串 s 切片
				path = append(path, s[start:i+1])
				// 递归：从下一个位置 i+1 开始探索
				backTracking(i + 1)

				// 回溯：清理现场
				path = path[:len(path)-1]
			}
		}
	}
	// 初始调用：从下标 0 开始
	backTracking(0)
	return res
}
