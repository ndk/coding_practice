package exercise

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	result := make([]int, m*n)
	p := 0

	var limit int
	if n < m {
		limit = n
	} else {
		limit = m
	}

	for depth := 0; 2*depth < limit; depth++ {
		// left to right
		for from, to := depth, m-depth; from < to; from++ {
			result[p] = matrix[depth][from]
			p++
		}
		// top to bottom
		for from, to, column := 1+depth, n-depth, m-1-depth; from < to; from++ {
			result[p] = matrix[from][column]
			p++
		}

		if p >= len(result) {
			break
		}

		// right to left
		for i, row := m-2-depth, n-1-depth; i >= depth; i-- {
			result[p] = matrix[row][i]
			p++
		}
		// bottom to top
		for i := n - 2 - depth; i > depth; i-- {
			result[p] = matrix[i][depth]
			p++
		}
	}

	return result
}
