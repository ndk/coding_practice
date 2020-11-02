package exercise

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([]bool, m*n)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == '1' && !visited[i*n+j] {
			visited[i*n+j] = true
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}

	components := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i*n+j] {
				components++
				dfs(i, j)
			}
		}
	}
	return components
}
