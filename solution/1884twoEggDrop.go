package solution

//dp 时间复杂度： O(2logn), n为楼层数 空间复杂度：O(2n)
//二维数组
// func twoEggDrop(n int) int {
// 	dp := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		dp[i] = make([]int, 3)
// 	}
// 	m := 0
// 	for dp[m][2] < n {
// 		m++
// 		for i := 2; i > 0; i-- {
// 			dp[m][i] = dp[m-1][i-1] + dp[m-1][i] + 1
// 		}
// 	}
// 	return m
// }

//dp 时间复杂度： O(2logn), n为楼层数 空间复杂度：O(2)
//一维数组优化，滚动数组的思想
func twoEggDrop(n int) int {
    ans, dp := 0, make([]int, 3)
    for dp[2]<n {
		ans++;
		for i := 2; i > 0; i-- {
			dp[i] += dp[i-1] + 1;
		}
	}
	return ans
}
