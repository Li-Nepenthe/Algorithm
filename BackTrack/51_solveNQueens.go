package BackTrack

func solveNQueens(n int) [][]string {
	col := make([]bool, n)
	// 主负对角线数量均为2n-1 但是编号从1开始
	diag1 := make([]bool, 2*n) // 主对角线
	diag2 := make([]bool, 2*n) // 副对角线
	queens := make([]int, n)   // queens[row] = col
	for i := range queens {
		queens[i] = -1
	}
	res := make([][]string, 0)
	var isValid func(int, int) bool
	isValid = func(i int, j int) bool {
		// 传入指定的坐标，检验所在列、主对角线、副对角线是否占用
		// 对于坐标(i,j) 所在列为j,所在主对角线为i-j+n 从1到2n-1,负对角线为 i+j 从0到2n-2
		// 如果三个均为占用，则表明当前可以放置皇后queen
		if !col[j] && !diag1[i-j+n] && !diag2[i+j] {
			queens[i] = j
			col[j], diag1[i-j+n], diag2[i+j] = true, true, true
			return true
		}
		return false
	}
	var backTracking func(int)
	backTracking = func(deep int) {
		// 能走到最后一层说明结果一定符合
		if deep == n {
			path := make([]string, 0)
			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}
			for i := 0; i < n; i++ {
				b[queens[i]] = 'Q'
				path = append(path, string(b))
				b[queens[i]] = '.'
			}
			cp := make([]string, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}
		for i := 0; i < n; i++ {
			//检验当前坐标是否能放置Q
			if isValid(deep, i) { //如果能进if循环，则状态已经被更改
				backTracking(deep + 1)
				//回溯
				queens[deep] = -1
				col[i], diag1[deep-i+n], diag2[deep+i] = false, false, false
			}
		}
	}
	backTracking(0)
	// 得到queens的存放结果后构造res结果集
	return res
}

func solveNQueensImprove(n int) [][]string {
	col := make([]bool, n)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)
	queens := make([]int, n)
	var res [][]string
	var backTracking func(int)
	backTracking = func(row int) {
		if row == n {
			path := make([]string, n)
			for i := 0; i < n; i++ {
				rowStr := make([]byte, n)
				for j := 0; j < n; j++ {
					if queens[i] == j {
						rowStr[j] = 'Q'
					} else {
						rowStr[j] = '.'
					}
				}
				path[i] = string(rowStr)
			}
			res = append(res, path)
			return
		}

		for j := 0; j < n; j++ {
			if !col[j] && !diag1[row-j+n] && !diag2[row+j] {
				queens[row] = j
				col[j], diag1[row-j+n], diag2[row+j] = true, true, true
				backTracking(row + 1)
				col[j], diag1[row-j+n], diag2[row+j] = false, false, false
			}
		}
	}

	backTracking(0)
	return res
}
