package Greedy

// 本题最优解思路
func canJump(nums []int) bool {
	reach := nums[0] + 0 // 当前能触及的范围即是第一个能走的范围
	for i, v := range nums {
		// 如果走到范围尽头，范围还没变大，则表示绝对走不出去
		if i > reach {
			return false
		}
		// 如果走到下一个，下一个能触及的范围比当前的范围还大则更新最大范围，否则最大范围不变
		// 这样能保证每次决策走的距离最远
		reach = max(reach, i+v)
		if reach >= len(nums)-1 {
			return true
		}
	}
	return false
}
