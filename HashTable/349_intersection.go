package HashTable

func intersection(nums1 []int, nums2 []int) []int {
	array := make([]int, 1001)
	for i := 0; i < len(nums1); i++ {
		array[nums1[i]]++
	}
	var r [1001]int
	count := 0
	for i := 0; i < len(nums2); i++ {
		if array[nums2[i]] > 0 {
			r[count] = nums2[i]
			array[nums2[i]] = 0
			count++
		}
	}
	return r[:count]
}
