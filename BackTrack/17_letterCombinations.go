package BackTrack

func letterCombinations(digits string) []string {
	mapping := [][]string{
		0: {""},            // 数字 0
		1: {""},            // 数字 1
		2: {"a", "b", "c"}, // 数字 2 直接对应索引 2
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	res := make([]string, 0)
	var path []byte
	var backTracking func(int)
	backTracking = func(index int) {
		// 如果index等于数组长度 表明处理完所有的数字了 保存结果并返回
		if index == len(digits) {
			res = append(res, string(path)) // 转换并存入结果
			return
		}

		digit := digits[index] - '0'
		//得到对应的数组
		letters := mapping[digit]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i][0]) // letter是[]string类型 letter[i]是string类型 letter[i][j]是取letter[i]中的byte字符 为byte类型
			backTracking(index + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}
