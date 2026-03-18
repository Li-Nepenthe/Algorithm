package Array

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	var n, m int
	if _, err := fmt.Fscan(reader, &n, &m); err != nil {
		return
	}

	// 2. 统计行总和与列总和
	rowSums := make([]int, n)
	colSums := make([]int, m)
	totalSum := 0

	// 计算每列/行的值
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var val int
			_, _ = fmt.Fscan(reader, &val)
			rowSums[i] += val
			colSums[j] += val
			totalSum += val
		}
	}

	// 3. 计算行/列的前缀和
	for i := 1; i < n; i++ {
		rowSums[i] += rowSums[i-1]
	}
	for j := 1; j < m; j++ {
		colSums[j] += colSums[j-1]
	}

	// 4. 计算最小差值
	minDiff := math.MaxInt64

	// 水平切割
	for i := 0; i < n-1; i++ {
		firstPart := rowSums[i]
		secondPart := totalSum - firstPart
		diff := abs(firstPart - secondPart)
		if diff < minDiff {
			minDiff = diff
		}
	}

	// 垂直切割
	for j := 0; j < m-1; j++ {
		firstPart := colSums[j]
		secondPart := totalSum - firstPart
		diff := abs(firstPart - secondPart)
		if diff < minDiff {
			minDiff = diff
		}
	}

	_, _ = fmt.Fprintln(writer, minDiff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
