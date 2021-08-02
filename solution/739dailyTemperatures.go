package solution

//单调栈 onon
func dailyTemperatuans(T []int) []int {
	ans := make([]int, len(T))
	var stack []int
	for i, t := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			index := stack[len(stack)-1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
