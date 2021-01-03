package exercise

func numOfPathsToDest(n int) int {
	if n == 1 {
		return 0
	}
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = 1
	}
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			row[j] += row[j-1]
		}
	}
	return row[n-1]
}
