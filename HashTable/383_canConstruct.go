package HashTable

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int, len(magazine))
	for i := range magazine {
		m[magazine[i]]++
	}

	for i := range ransomNote {
		if value, ok := m[ransomNote[i]]; ok && value > 0 {
			m[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}

// 可以不适用map 由于所有可能都是由小写字母组成 可以定义一个[26]int

func canConstruct1(ransomNote string, magazine string) bool {
	array := make([]int, 26)
	for i := range magazine {
		array[magazine[i]-'a']++
	}
	for i := range ransomNote {
		if array[ransomNote[i]-'a'] > 0 {
			array[ransomNote[i]-'a']--
		} else {
			return false
		}
	}
	return true
}
