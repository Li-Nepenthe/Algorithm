package Array

//声明二维数组

func makeMatrix(length int) [][]int {
	//动态分配
	matrix1 := make([][]int, length) // 创建了一个长度为length的切片其中每个元素都为[]int切片 行指针
	for i := range matrix1 {
		matrix1[i] = make([]int, length)
	}

	//固定初始长度
	var _ [5][5]int

	//字面量初始化
	var _ = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	return matrix1
}

func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	var top = 0        //上边界
	var right = n - 1  //右边界
	var bottom = n - 1 //下边界
	var left = 0       //左边界
	var currentNum = 1
	for currentNum <= n*n {
		// 从左到右填充top行 填充完 top++   从左边界到右边界
		for i := left; i <= right; i++ {
			matrix[top][i] = currentNum // 列增加 行不变
			currentNum++
		}
		top++
		// 从上到下填充right列，填充完 right--  从上边界到下边界
		for i := top; i <= bottom; i++ {
			matrix[i][right] = currentNum // 行增加 列不变
			currentNum++
		}
		right--
		// 从右到左填充bottom行，填充完，bottom--  从右边界到左边界
		for i := right; i >= left; i-- {
			matrix[bottom][i] = currentNum // 列减少，行不变
			currentNum++
		}
		bottom--
		// 从下到上填充left列，填充完 left--   从下边界到上边界
		for i := bottom; i >= top; i-- {
			matrix[i][left] = currentNum // 行减少，列不变
			currentNum++
		}
		left++
	}
	return matrix
}
