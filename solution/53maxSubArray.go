package solution

//动态规划 on o1
func maxSubArray(nums []int) int {
	n, ans := len(nums), nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
