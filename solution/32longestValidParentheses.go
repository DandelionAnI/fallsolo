package solution

//1.栈 on on
func longestValidParentheses(s string) int {
	stack, ans := []int{}, 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }