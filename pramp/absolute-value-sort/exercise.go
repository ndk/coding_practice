package exercise

import (
	"math"
	"sort"
)

func absSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		if math.Abs(float64(arr[i])) < math.Abs(float64(arr[j])) {
			return true
		}
		if math.Abs(float64(arr[i])) > math.Abs(float64(arr[j])) {
			return false
		}
		return arr[i] < arr[j]
	})
}
