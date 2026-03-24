package HashTable

import "sort"

// 此方法虽然能得到期望的数组 但是无法去重
//func threeSum(nums []int) [][]int {
//
//	var res [][]int
//	// 先统计数组中每个数出现的次数，避免重复
//	m := make(map[int]int, len(nums))
//	for i := range nums {
//		m[nums[i]]++
//	}
//	// 不断求和得a+b 然后寻找map里是否有对应的c = -(a+b)
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			a, b := nums[i], nums[j]
//			sum := a + b
//			exp := -sum              //期望值
//			if _, ok := m[exp]; ok { // 如果确实出现在map 判断该值是否利用重复
//				// 检验是否重复
//				m[a]--
//				m[b]--
//				m[exp]--
//				if m[a] >= 0 && m[b] >= 0 && m[exp] >= 0 {
//					// 如何获得exp的角标？
//					res = append(res, []int{a, b, exp})
//				}
//					m[a]++
//					m[b]++
//					m[exp]++
//				}
//			}
//		}
//	}
//
//	return res
//}

// 思路: 既然数组杂乱无序，无法分辨是否重复，不如直接先进行排序 然后再遍历
// 如nums = [-1,0,1,2,-1,-4] 排序后nums = [-4,-1,-1,0,1,2] 前面的-1 不会对后面的-1产生影响
// 这道题要对 i,j,k三个元素都要去重
func threeSum(nums []int) [][]int {
	// 虽然题目说明len >= 3 但是还是做防御性编程
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	// 首先对数组元素进行排序，因为题目只要求返回三元组的值而不是下标
	sort.Ints(nums)
	// 如果排序后第一个数都大于0 则不可能存在三元组为0
	if nums[0] > 0 {
		return nil
	}
	// 接下来固定i，然后去寻找j和k 看看能否满足nums[i]+nums[j]+nums[k] == 0
	for i := 0; i < len(nums)-2; i++ { // i遍历到倒数第三个即可
		// 首先针对i去重 这里注意要和之前的比而不是和之后的比， 比如{-1，-1，2}
		// 比如i遍历到第二个-1 和前一个相比发现相同 于是便跳过
		// 如果写nums[i] == nums[i+1] 数组内去重，会略掉许多三元组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right { //边界条件 left right不能相等 否则不存在三个不同元素
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 { //说明三元素之和过大 right左移
				right--
			} else if sum < 0 {
				left++
			} else { //nums[i]+nums[j]+nums[k] == 0
				// 这里要对j和k进行去重操作 j和k要跳到下一个和当前元素不一样的位置 如{0,0,0,0,0}
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//首先利用for循环找到最后一个满足当前位置的值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 找到位置后 left和right还得移动
				left++
				right--
			}
		}
	}
	return res
}
