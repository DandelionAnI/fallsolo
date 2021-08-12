package solution

//模拟 om+n o1
func findNumberIn2DArray(matrix [][]int, target int) bool {
		if len(matrix) == 0 {
			return false
		}
		i, j := 0, len(matrix[0])-1
		for j >= 0 && i <= len(matrix)-1 {
			if target == matrix[i][j] {
				return true
			} else if target > matrix[i][j] {
				i++
			} else {
				j--
			}
		}
		return false
}
