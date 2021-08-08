package solution

//异或，找结果第一个为1的标志位，分组再异或，on o1
func singleNumber2(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	p := 1
	for sum&p == 0 {
		p <<= 1
	}
	nums1, nums2 := 0, 0
	for _, v := range nums {
		if v&p == 0 {
			nums1 ^= v
		} else {
			nums2 ^= v
		}
	}
	return []int{nums1, nums2}
}
