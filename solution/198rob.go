package solution

//dp on on
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// //dp优化 on o1
// func rob_1(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	cur, pre, tmp := 0, 0, 0
// 	for i := 0; i < n; i++ {
// 		tmp = cur
// 		cur = max(cur, pre+nums[i])
// 		pre = tmp
// 	}
// 	return cur
// }
