package solution

func longestPalindrome(s string) string {
	n := len(s)
	ans, dp := "", make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (ans == "" || j-i+1 > len(ans)) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
