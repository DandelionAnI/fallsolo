package solution

//双指针 on o1
func nextPermutation(nums []int) {
	n := len(nums) - 1
	l := n - 1
	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}
	if l >= 0 {
		r := n
		for r >= 0 && nums[l] >= nums[r] {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	reverse(nums[l+1:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
