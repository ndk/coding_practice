package exercise

import "sort"

func findPairsWithGivenDifference(arr []int, k int) [][2]int {
	index := make([]int, len(arr))
	for i := range index {
		index[i] = i
	}

	sort.Slice(index, func(i, j int) bool {
		return arr[index[i]] < arr[index[j]]
	})

	var result [][2]int
	for p1, p2 := 0, 1; p2 < len(index); {
		diff := arr[index[p2]] - arr[index[p1]]
		if diff == k {
			result = append(result, [2]int{index[p2], index[p1]})
			p1++
			p2++
		} else if diff < k {
			p2++
		} else {
			p1++
		}
	}

	if len(result) == 0 {
		return nil
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][1] < result[j][1]
	})

	result2 := make([][2]int, len(result))
	for i, pair := range result {
		result2[i] = [2]int{arr[pair[0]], arr[pair[1]]}
	}

	return result2
}

