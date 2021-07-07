package solution

// var dir = [][]int{
// 	{-1, 0},
// 	{0, 1},
// 	{1, 0},
// 	{0, -1},
// }

//dfs omn omn
func numIslands(grid [][]byte) int {
	h, w := len(grid), len(grid[0])
	var count func(x, y int)
	count = func(x, y int) {
		grid[x][y] = '0'
		for i := 0; i < 4; i++ {
			newx := x + dir[i][0]
			newy := y + dir[i][1]
			if newx >= 0 && newx < h && newy >= 0 && newy < w {
				if grid[newx][newy] == '1' {
					count(newx, newy)
				}
			}
		}
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '1' {
				count(i, j)
				ans++
			}
		}
	}
	return ans
}
