package solution

func permutation(s string) []string {
	ans := []string{}
	//判断当前字符串是否在结果集中
	ifAns := map[string]bool{}
	//判断当前字符有无被访问
	ifCheck := map[int]bool{}
	var dfs func(meta string)
	dfs = func(meta string) {
		if len(meta) == len(s) {
			if ifAns[meta] == false {
				ans = append(ans, meta)
				ifAns[meta] = true
			}
			return
		}
		for i := 0; i < len(s); i++ {
			if ifCheck[i] {
				continue
			}
			ifCheck[i] = true
			dfs(meta + string(s[i]))
			ifCheck[i] = false
		}
	}
	dfs("")
	return ans
}
