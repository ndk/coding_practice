package exercise

import (
	"strconv"
	"strings"
)

func dfs(ans *[]string, seen map[string]struct{}, node string, k int) {
	for x := 0; x < k; x++ {
		nei := node + strconv.Itoa(x)
		if _, ok := seen[nei]; !ok {
			seen[nei] = struct{}{}
			dfs(ans, seen, nei[1:], k)
			*ans = append(*ans, strconv.Itoa(x))
		}
	}
}

func crackSafe(n int, k int) string {
	if n == 1 && k == 1 {
		return "0"
	}

	start := strings.Repeat("0", n-1)

	seen := map[string]struct{}{}
	var ans []string
	dfs(&ans, seen, start, k)
	ans = append(ans, start)
	return strings.Join(ans, "")
}
