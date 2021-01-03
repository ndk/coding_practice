package exercise

func uniquePaths(m int, n int) int {
	if m == 1 {
		return 1
	}
	row := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] += row[j-1]
		}
	}
	return row[n-1]
}

///////////////////////////////////////////////////////////////////////////////
// From Leetcode
// Approach 2: Math

func uniquePaths2(m int, n int) int {
	d := m
	if n > m {
		d = n
	}

	t := n + m - 1

	result := 1
	for i, to := 1, t-d+1; i < to; i++ {
		result *= t - i
		result /= i
	}

	return result
}

func uniquePaths3(m int, n int) int {
	m, n = m-1, n-1
	if m == 0 || n == 0 {
		return 1
	}
	// ensure m is always greater,
	// otherwise some test cases result in overflow
	if m < n {
		m, n = n, m
	}

	d1, d2 := 1, 1
	for i, to := m+1, m+n; i <= to; i++ {
		d1 *= i
	}
	for i := 1; i <= n; i++ {
		d2 *= i
	}
	return d1 / d2
}
