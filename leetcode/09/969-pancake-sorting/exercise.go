package exercise

func pancakeSort(arr []int) []int {
	var result []int
	flip := func(k int) {
		if k > 0 {
			result = append(result, k+1)
		}
		for l := 0; l < k; {
			arr[l], arr[k] = arr[k], arr[l]
			l++
			k--
		}
	}

	for unsortedRange := len(arr) - 1; unsortedRange > 0; unsortedRange-- {
		largest := 0
		for i := 1; i <= unsortedRange; i++ {
			if arr[largest] < arr[i] {
				largest = i
			}
		}
		flip(largest)
		flip(unsortedRange)
	}

	return result
}
