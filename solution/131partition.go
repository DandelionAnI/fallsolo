package solution

//dp+dfs on2^n on^2
func partition1(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	check := make([][]bool, n)
	for i := 0; i < n; i++ {
		check[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			check[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			check[i][j] = (s[i] == s[j] && check[i+1][j-1])
		}
	}

	str := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			tmp := []string{}
            tmp = append(tmp, str...)
            ans = append(ans, tmp)
			return
		}
		for j := i; j < n; j++ {
			if check[i][j] {
				str = append(str, s[i:j+1])
				dfs(j+1)
				str = str[:len(str)-1]
			}
		}
	}
	dfs(0)
	return ans
}
