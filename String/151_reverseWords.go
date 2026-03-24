package String

// 利用快慢指针，不断地把字符往前挪动 然后对整体字符反转 在对单独的单词反转
// 例：__abc__def_ --> abc_def -->fed_cba-->def_abc

func reverseWords(s string) string {

	//将字符串转换为byte数组
	bytes := []byte(s)
	length := len(bytes)
	fast, slow := 0, 0

	for fast < length && bytes[fast] == ' ' {
		fast++
	}
	for fast < length {
		if slow != 0 { //当前slow指的不是第一个单词
			bytes[slow] = ' '
			slow++
		}

		// 当前fast所指为字符 与 slow所在位置交换
		// 情况一：无前导空格，fast slow同步移动 直到fast指向空格
		// 情况二：有前导空格，当前单词整体偏移
		for fast < length && bytes[fast] != ' ' {
			bytes[slow] = bytes[fast]
			slow++
			fast++
		}
		// 扫描完一个单词 让fast指向下一个单词 如果是尾随空格则全部跳过
		for fast < length && bytes[fast] == ' ' {
			fast++
		}
		//此时fast指向了新的单词第一个（如果有）
	}
	bytes = bytes[:slow]

	// 交换整个数组
	swapByte(bytes)
	// 现在交换单词顺序，记录非空格的起始位置和结束位置
	start, end := 0, 0
	for end != len(bytes) {
		if end == len(bytes)-1 {
			swapByte(bytes[start : end+1])
		}
		if bytes[end] != ' ' {
			end++
		} else {
			swapByte(bytes[start:end])
			end = end + 1
			start = end
		}
	}
	return string(bytes)

}

func swapByte(bytes []byte) {
	length := len(bytes)
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
}
