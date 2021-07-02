package solution

//回溯 dfs on o2n
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	findGenerateParenthesis(n, n, "", &ans)
	return ans
}

func findGenerateParenthesis(left, right int, str string, ans *[]string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, str)
		return
	}
	if left > 0 {
		findGenerateParenthesis(left-1, right, str+"(", ans)
	}

	if right > 0 && left < right {
		findGenerateParenthesis(left, right-1, str+")", ans)
	}
}
