package exercise

import (
	"sort"
)

func findGrantsCap(grantsArray []int, newBudget int) float32 {
	n := len(grantsArray)
	if n == 0 {
		return 0
	}

	sort.Ints(grantsArray)
	budget, prev := 0, 0
	for i, grant := range grantsArray {
		right := (grant - prev) * (n - i)
		if right > newBudget-budget {
			return float32(prev) + float32(newBudget-budget)/float32(n-i)
		}
		budget += right
		prev = grant
	}
	return float32(grantsArray[n-1])
}
