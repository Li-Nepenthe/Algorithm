package String

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	// k可能的长度
	for k := 1; k <= length/2; k++ {
		// 如果字符串长度不是当前可能字串长度的整数倍 则当前的k一定不满足
		if length%k != 0 {
			continue
		}
		// 如果前k个字符是模板，那么第j个字符必须等于j-k个字符
		for j := k; j < length; j++ {
			// 如果不满足 说明不是子串 跳过该k
			if s[j] != s[j-k] {
				break
			}
			if j == length-1 {
				return true
			}
		}
	}
	return false
}
