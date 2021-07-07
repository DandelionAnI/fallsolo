package solution

//回溯 dfs o(MN3^L) omn
func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	//作原图的bool二维数组记录走过的路
	wordmap := make([][]bool, h)
	for i := range wordmap {
		wordmap[i] = make([]bool, w)
	}
	var check func(x, y, k int) bool
	check = func(x, y, k int) bool {
		//当前字符不匹配退出，剪枝
		if board[x][y] != word[k] {
			return false
		}
		// 单词存在于网格中
		if k == len(word)-1 {
			return true
		}
		wordmap[x][y] = true
		// 回溯时已访问的单元格还原
		defer func() { wordmap[x][y] = false }()
		// 走下一步并回溯
		for i := 0; i < 4; i++ {
			xnew := x + dir[i][0]
			ynew := y + dir[i][1]
			if xnew >= 0 && xnew < h && ynew >= 0 && ynew < w && !wordmap[xnew][ynew] {
				if check(xnew, ynew, k+1) {
					return true
				}
			}
		}
		return false
	}
	for x, row := range board {
		for y := range row {
			if check(x, y, 0) {
				return true
			}
		}
	}
	return false
}
