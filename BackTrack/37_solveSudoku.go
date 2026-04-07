package BackTrack

// 数独和N皇后的区别在于，N皇后每行放置一个后就可以递归到下一层
// 但是数独中每一行有多个空格需要进行尝试，无法在该行填入一个空格就递归到下一行
// 最初的想法是把每一行所有的情况都列举出来完毕后在进行下一行，结果控制十分复杂，既要向右走又要向下走
// 所以逻辑转换，不在按行或者按列递归，而是按照空格的顺序一次递归（行优先）
// 首先遇到第一行第一个空格，尝试填入数字校验合法性 然后列移动到下一个空格，知道某一个行空格全部递归完 然后行+1
// 同时要利用函数的返回值，如果递归到最深层，则表明之前所有逻辑全部通过，当前棋盘已经处理完毕，所以要层层返回
// 这个传递的顺序其实就是空格的递归顺序，也是回溯的调用顺序
// 比如倒数第二个空格能否成功，依赖最后一个能否成功，如果最后一个返回成功，则倒数第二个也会以此往上返回成功，状态传递
// 如果全部都是返回成功则所有针对board的回溯代码都不会执行，board被原样保留
// 如果哪一层失败，则在该层回溯
// 所以最后要么返回true 得到一个合法board，要么终止递归 返回false

func solveSudoku(board [][]byte) {
	var isValid func(int, int, byte) bool
	isValid = func(row int, col int, target byte) bool {
		//将target转为字节然后进行比较
		//判断对应列
		for i := 0; i < 9; i++ {
			if board[i][col] != '.' && board[i][col] == target {
				return false
			}
		}
		// 接着判断对应九宫格是否出现重复
		// 1.先利用row判断在所在九宫格是行第几个 再用列判断所在九宫格在列第几个
		rIndex, cIndex := (row/3)*3, (col/3)*3
		// 2.定位到九宫格后，遍历九宫格
		for i := rIndex; i < rIndex+3; i++ {
			for j := cIndex; j < cIndex+3; j++ {
				if board[i][j] != '.' && board[i][j] == target {
					return false
				}
			}
		}
		return true
	}
	var backTracking func(r, c int) bool
	backTracking = func(r, c int) bool {
		// 如果当前的列已经走到投了 直接换下一行
		if c == 9 {
			return backTracking(r+1, 0)
		}
		// 如果行已经走到头了 说明棋盘已经完成 直接返回成功状态
		if r == 9 {
			return true
		}
		// 检验当前的格子是否需要填值，如果不需要跳到下一个
		if board[r][c] != '.' {
			return backTracking(r, c+1)
		}
		// 现在处理当前所在的格子需要填值的情况

		// 先遍历出当前行已有的数字 避免重复选取
		used := [10]bool{} // 记录当前行 row 已有的数字
		for j := 0; j < 9; j++ {
			if board[r][j] != '.' {
				used[board[r][j]-'0'] = true
			}
		}

		// 尝试填入时，利用used数组跳过已有的数
		for num := 1; num <= 9; num++ {
			// 如果行内已经有了，直接跳过
			if used[num] {
				continue
			}
			// 对可能选取的数进行合法性判断，合法才递归否则继续跳过
			if isValid(r, c, byte(num+'0')) {
				board[r][c] = byte(num + '0')
				// 如果传递到最后的逻辑 返回true 则利用此处层层返回 不再修改board的值
				if backTracking(r, c+1) {
					return true
				}
				board[r][c] = '.'
			}
		}
		return false
	}
	backTracking(0, 0)
}
