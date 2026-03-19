package HashTable

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var array = make([]int, 26)
	for i := 0; i < len(s); i++ {
		array[s[i]-'a']++
		array[t[i]-'a']--
	}
	for i := 0; i < len(s); i++ {
		if array[s[i]-'a'] != 0 {
			return false
		}
	}
	return true
}
