package exercise

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return nil
	}

	sort.Ints(arr)

	result := [][]int{{arr[0], arr[1]}}
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < min {
			min = diff
			result = [][]int{{arr[i-1], arr[i]}}
		} else if diff == min {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}

	return result
}
