package BackTrack

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	length := len(s)
	if length < 4 || length > 12 {
		return []string{}
	}
	res := make([]string, 0)
	var path []string
	var isValid func(s string) bool
	isValid = func(s string) bool {
		if len(s) == 0 {
			return false
		}
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		number, _ := strconv.Atoi(s)
		return number >= 0 && number <= 255
	}
	var backTracking func(string, int)
	backTracking = func(st string, length int) {
		if len(st) == 0 && length == 0 {
			res = append(res, strings.Join(path, "."))
			return
		}
		if length == 0 {
			return
		}
		for i := 0; i < len(st); i++ {
			if isValid(st[:i+1]) {
				path = append(path, st[:i+1])
				backTracking(st[i+1:], length-1)
				//回溯
				path = path[:len(path)-1]
			}
			continue
		}
	}
	backTracking(s, 4)
	return res
}
