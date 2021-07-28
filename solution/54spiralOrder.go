package solution

//螺旋遍历 omn o1
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	index := len(matrix) * len(matrix[0])

	for index >= 1 {
		for i := l; i <= r && index >= 1; i++ {
			ans = append(ans, matrix[t][i])
			index--
		}
		t++
		for i := t; i <= b && index >= 1; i++ {
			ans = append(ans, matrix[i][r])
			index--
		}
		r--
		for i := r; i >= l && index >= 1; i-- {
			ans = append(ans, matrix[b][i])
			index--
		}
		b--
		for i := b; i >= t && index >= 1; i-- {
			ans = append(ans, matrix[i][l])
			index--
		}
		l++
	}
	return ans
}
