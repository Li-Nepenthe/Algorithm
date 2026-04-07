package Greedy

import (
	"cmp"
	"fmt"
	"slices"
)

// 排序背后的底层逻辑

// slices.SortFunc是排序的执行者，它传入两个参数，一个是待排序对象，一个是如何排序的逻辑
// 如何自定义排序的核心就是如何自定义排序的逻辑
// slices.SortFunc的接受的排序逻辑是这样的：抽象理解为每次拿到两个值a，b
// 然后比较a和b的大小，如果返回值为1，则把a放在b后面 -1则把a放在b前面 0则可以表示不动
// 如果需要逆序只需修改逻辑或者更改传入顺序，如我定义a<b反而传回1 这样a比b小时 a反而在b后面
// 而compare的逻辑就是如此 他会比较传入的两个参数大小 并返回-1 0 1三个值
// 针对多重逻辑排序如二维数组情况

// Sort1D 1. 一维切片排序 (原生类型)
func Sort1D() {
	// 1.1 整型 (升序 & 降序)
	ints := []int{5, 2, 9, 1, 5}
	// ！！！ 针对原生数据类型的升序排列 推荐使用slices.Sort 如果降序和原生数据类型以及需要定制排序的可以使用slices.SortFunc
	slices.Sort(ints)
	slices.SortFunc(ints, func(a, b int) int { return cmp.Compare(a, b) }) // 升序: 1, 2, 5, 5, 9
	slices.SortFunc(ints, func(a, b int) int { return cmp.Compare(b, a) }) // 降序: 9, 5, 5, 2, 1

	// 1.2 浮点型 (安全处理 NaN 和 -0.0)
	floats := []float64{3.14, 1.59, 2.65}
	slices.SortFunc(floats, func(a, b float64) int { return cmp.Compare(a, b) }) // 升序

	// 1.3 字符串 (按字典序)
	strs := []string{"banana", "apple", "cherry"}
	slices.SortFunc(strs, func(a, b string) int { return cmp.Compare(b, a) }) // 降序: cherry, banana, apple
}

// Sort2D 2. 二维切片排序 (数组的数组)
func Sort2D() {
	// 以区间为例: [起点, 终点]
	intervals := [][]int{{1, 5}, {2, 6}, {1, 3}, {8, 10}, {2, 4}}

	// 2.1 只看第二个元素 (升序)
	slices.SortFunc(intervals, func(a, b []int) int {
		// 每次拿到两个值如{1, 5} {2, 4} 只比较第二个元素的情况 4 < 5 所以 {2, 4}在{1, 5}前面
		return cmp.Compare(a[1], b[1])
	})

	// 2.2 先看第一个(升序)，再看第二个(升序)
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(a[0], b[0]),
			cmp.Compare(a[1], b[1]),
		)
	})

	// 2.3 先看第二个(降序)，再看第一个(升序) - 常见于某些贪心策略
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Or(
			cmp.Compare(b[1], a[1]), // 第二个元素降序 (b, a)
			cmp.Compare(a[0], b[0]), // 第一个元素升序 (a, b)
		)
	})

	// 2.4 字符串二维数组：按第一列字符串长度排，长度一样按字典序
	strMatrix := [][]string{{"abc", "z"}, {"a", "x"}, {"abcd", "y"}}
	slices.SortFunc(strMatrix, func(a, b []string) int {
		return cmp.Or(
			cmp.Compare(len(a[0]), len(b[0])), // 比较长度
			cmp.Compare(a[0], b[0]),           // 比较字典序
		)
	})
}

// Student 3. 自定义结构体排序 (业务开发最常用)
type Student struct {
	ID    int
	Name  string
	Score float64
}

func SortStructs() {
	// ⚠️ 最佳实践：结构体较大时，切片里存指针，避免排序时发生海量内存拷贝
	students := []*Student{
		{ID: 2, Name: "Bob", Score: 95.5},
		{ID: 1, Name: "Alice", Score: 95.5},
		{ID: 3, Name: "Charlie", Score: 88.0},
	}

	// 3.1 单条件：按 Score 升序
	slices.SortFunc(students, func(a, b *Student) int {
		return cmp.Compare(a.Score, b.Score)
	})

	// 3.2 多条件组合：先按 Score 降序，Score 相同按 ID 升序
	slices.SortFunc(students, func(a, b *Student) int {
		return cmp.Or(
			cmp.Compare(b.Score, a.Score), // 分数降序
			cmp.Compare(a.ID, b.ID),       // ID升序
		)
	})
}

func main() {
	// 你可以在这里调用上面的函数进行打印测试
	Sort1D()
	Sort2D()
	SortStructs()
	fmt.Println("排序档案加载完毕！")
}
