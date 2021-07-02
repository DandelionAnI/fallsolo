package solution

//异或 on o1
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
