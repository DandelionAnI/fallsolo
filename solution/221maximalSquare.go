package solution

//dp omn omn(可以用两个一维的数组，可以优化到on)
func maximalSquare(matrix [][]byte) int {
	n, m := len(matrix),len(matrix[0])
	dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }
    maxSide := 0
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
                    maxSide = dp[i][j]
                }
            }
        }
    }
    return maxSide * maxSide
}

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
