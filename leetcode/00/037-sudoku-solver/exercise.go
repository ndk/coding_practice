package exercise

func solveSudoku(board [][]byte) {
	rowSeen, columnSeen, subSeen := [9][9]bool{}, [9][9]bool{}, [3][3][9]bool{}
	for r := 0; r < 9; r++ {
		sr := r / 3
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				i := board[r][c] - '1'
				rowSeen[r][i], columnSeen[c][i], subSeen[sr][c/3][i] = true, true, true
			}
		}
	}

	var recursive func(board [][]byte, r, c int, rowSeen *[9][9]bool, columnSeen *[9][9]bool, subSeen *[3][3][9]bool) bool
	recursive = func(board [][]byte, r, c int, rowSeen *[9][9]bool, columnSeen *[9][9]bool, subSeen *[3][3][9]bool) bool {
		for {
			if c == 9 {
				c = 0
				r++
			}
			if r == 9 {
				return true
			}
			if board[r][c] == '.' {
				break
			}
			c++
		}

		sr, sc := r/3, c/3
		for i := 0; i < 9; i++ {
			if rowSeen[r][i] || columnSeen[c][i] || subSeen[sr][sc][i] {
				continue
			}
			rowSeen[r][i], columnSeen[c][i], subSeen[sr][sc][i] = true, true, true
			board[r][c] = '1' + byte(i)
			if recursive(board, r, c+1, rowSeen, columnSeen, subSeen) {
				return true
			}
			rowSeen[r][i], columnSeen[c][i], subSeen[sr][sc][i] = false, false, false
			board[r][c] = '.'
		}
		return false
	}

	recursive(board, 0, 0, &rowSeen, &columnSeen, &subSeen)
}
