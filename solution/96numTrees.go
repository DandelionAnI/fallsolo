package solution

//dp  on o1
//[1,2,3,4,…, i ,…,n-1,n]，以 i 为 根节点，那么左半边 [1,2,3,……,i-1] 和 
//右半边 [i+1,i+2,……,n-1,n] 分别能组成二叉排序树的不同个数相乘，
//即为以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数，也即 F(i,n)。
//所以状态转移方程是 dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]，
//最终要求的结果是 dp[n] 。
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
