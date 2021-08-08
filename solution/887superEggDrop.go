package solution

// dp 时间复杂度： O(klogn), n为楼层数 空间复杂度：O(kn)
// 二维数组
// func superEggDrop(k int, n int) int {
// 	dp := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		dp[i] = make([]int,k+1)
// 	}
// 	m := 0
// 	for dp[m][k] < n {
// 		m++
// 		for i := k; i > 0; i-- {
// 			dp[m][i] = dp[m-1][i-1] + dp[m-1][i] + 1
// 		}
// 	}
// 	return m
// }

//dp 时间复杂度： O(klogn), n为楼层数 空间复杂度：O(k)
//一维数组优化，滚动数组的思想
func superEggDrop(k int, n int) int {
    ans, dp := 0, make([]int, k+1)
    for dp[k]<n {
		ans++;
		for i := k; i > 0; i-- {
			dp[i] += dp[i-1] + 1;
		}
	}
	return ans
}