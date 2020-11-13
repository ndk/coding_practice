package exercise

import (
	"math"
	"strconv"
	"strings"
)

// Inverse Burrows-Wheeler Transform
func crackSafe_IBWT(n int, k int) string {
	m := int(math.Pow(float64(k), float64(n-1)))
	p := make([]int, m*k)
	for i := 0; i < k; i++ {
		for q := 0; q < m; q++ {
			p[i*m+q] = q*k + i
		}
	}

	var ans []string
	for i := 0; i < m*k; i++ {
		j := i
		for p[j] >= 0 {
			ans = append(ans, strconv.Itoa(j/m))
			v := p[j]
			p[j] = -1
			j = v
		}
	}

	for i := 0; i < n-1; i++ {
		ans = append(ans, "0")
	}
	return strings.Join(ans, "")
}
