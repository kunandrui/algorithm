package exist

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	wordLen := len(word)

	var dfs func(x, y, i int) bool
	dfs = func(x, y, i int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[i] {
			return false
		}
		if i == wordLen-1 {
			return true
		}
		board[x][y] = ' '
		res := dfs(x-1, y, i+1) || dfs(x, y+1, i+1) || dfs(x+1, y, i+1) || dfs(x, y-1, i+1)
		board[x][y] = word[i]
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
