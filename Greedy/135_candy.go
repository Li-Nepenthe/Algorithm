package Greedy

//如果一个孩子 左右两边的rating都大于等于他，那么他只能拿到最低的一个糖果

func candy(ratings []int) int {
	// 当前只有一个孩子 直接返回一个糖果
	if len(ratings) == 1 {
		return 1
	}
	candies := make([]int, len(ratings))
	candies[0] = 1
	// 先从左到右和左邻居比 再从右到左和右邻居比
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	sum := 0
	for i := 0; i < len(candies); i++ {
		sum += candies[i]
	}
	return sum
}
