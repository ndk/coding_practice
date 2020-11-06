package exercise

//   ^
//   |
//   |
// m |
//   |
//   |----------------->
//           n

func longestIncreasingPath(matrix [][]int) int {
	directions := [5]int{0, -1, 0, 1, 0}

	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	columns := len(matrix[0])

	cache := make([]int, rows*columns)

	var dfs func(row, column int) int
	dfs = func(row, column int) int {
		columns := len(matrix[0])

		pos := row*columns + column
		if cache[pos] != 0 {
			return cache[pos]
		}

		cache[pos] = 0
		for k := 0; k < 4; k++ {
			nextRow, nextColumn := directions[k]+row, directions[k+1]+column
			if nextRow < 0 || len(matrix) <= nextRow || nextColumn < 0 || len(matrix[row]) <= nextColumn {
				continue
			}
			if matrix[nextRow][nextColumn] <= matrix[row][column] {
				continue
			}
			depth := cache[nextRow*columns+nextColumn]
			if depth == 0 {
				depth = dfs(nextRow, nextColumn)
			}
			if depth > cache[pos] {
				cache[pos] = depth
			}
		}
		cache[pos]++
		return cache[pos]
	}

	longest := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if depth := dfs(r, c); depth > longest {
				longest = depth
			}
		}
	}

	return longest
}
