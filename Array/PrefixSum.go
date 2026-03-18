package Array

import "fmt"

func PrefixSum() int {
	var n int
	_, _ = fmt.Scan(&n)
	var array = make([]int, n)
	for i := range array {
		array[i], _ = fmt.Scan(&array[i])
	}
	prefixSum := make([]int, n+1)
	// prefixSum[0] 默认为 0，所以循环从 1 开始
	for i := 1; i <= n; i++ {
		// 当前“前 i 个数之和” = 前 i-1 个数之和 + 当前第 i 个数
		prefixSum[i] = prefixSum[i-1] + array[i-1]
	}
	var a, b int
	_, _ = fmt.Scan(&a)
	_, _ = fmt.Scan(&b)
	return prefixSum[b+1] - prefixSum[a]
}

func main() {
	fmt.Println(PrefixSum())
}
