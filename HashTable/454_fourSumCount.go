package HashTable

// nums1和nums2的各项之和 结果为a+b
// nums3和nums4的各项之和 结果为c+d
// 将a+b存入map 遍历nums3和nums4 遍历时同步计算c+d 并查找map中是否有a+b = -(c+d)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, len(nums1)*len(nums2))
	count := 0
	for i := range nums1 {
		for j := range nums2 {
			sum := nums1[i] + nums2[j]
			m[sum]++
		}
	}

	for i := range nums3 {
		for j := range nums4 {
			sum := nums3[i] + nums4[j] //当前元素的和
			if _, ok := m[-sum]; ok {  // 要找到有没有-sum
				count += m[-sum] // 当m[-sum]存在时其出现次数全部能满足
			}
		}
	}
	return count
}
