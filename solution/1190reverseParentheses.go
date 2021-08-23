package solution

//栈+遍历	on on
func reverseParentheses(s string) string {
	pair, stack := make([]int, len(s)), []int{}
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[pre] = pre, i
		}
	}
	ans := []byte{}
	for i, step := 0, 1; i < len(s); i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
