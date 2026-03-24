package String

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	// 遍历所有可能的起始位置
	// 注意：如果 n < m，i <= n-m 将为假，直接跳过循环返回 -1
	for i := 0; i <= n-m; i++ {
		// 利用切片直接比较
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}
