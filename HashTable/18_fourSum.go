package HashTable

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res = make([][]int, 0, 100)
	//状态剪枝

	// 如果数组长度小于4 直接返回
	if n < 4 {
		return nil
	}
	// 如果数组最小值的和都大于target 则退出循环
	if nums[0]+nums[1]+nums[2]+nums[3] > target {
		return nil
	}
	// 如果数组最大值的和都小于target 则退出循环
	if nums[n-1]+nums[n-2]+nums[n-3]+nums[n-4] < target {
		return nil
	}

	for i := 0; i < n-3; i++ { // i走到倒数第四个
		// 最外层去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 状态剪枝

		// 如果当前能拿到的最小和都大于target 那么后续的i开头只会更大 那么直接返回
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		// 如果当前的能拿到的最大值都小于target 那么得换一个i
		if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}

		// 内层还是套用3Sum的逻辑
		for j := i + 1; j < n-2; j++ { //j可以走到倒数第三个
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 状态剪枝
			// 此时i和j固定下来了，那么最小的数为下表为j和j+1，最大的两个数下标为N-2，N-1
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				// 那么当前的j已经大于target 换后续的j只会更大，所以退出j层循环
				break //此break就会传到到 if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target 导致整个循环结束
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
