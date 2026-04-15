package Greedy

func canCompleteCircuit(gas []int, cost []int) int {

	totalSum, curSum, index := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		//当前经过站所花的总支出
		totalSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		// 如果到达当前站点后油量小于等于0  即不足以开到下一个站点，那么当前的节点一定不行
		if curSum < 0 {
			//当前剩余油量归零 从下一站从新开始
			curSum = 0
			index = i + 1
		}
	}
	// 如果最后净支出大于0的话 则表示无法跑一整圈（因为一开始车就没有油）
	if totalSum < 0 {
		return -1
	}
	return index
}
