package exercise

func flip(arr []int, k int) {
	for l, r := 0, k-1; l < r; {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func pancakeSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		m := 0
		for j := 1; j <= i; j++ {
			if arr[j] > arr[m] {
				m = j
			}
		}
		flip(arr, m+1)
		flip(arr, i+1)
	}
}
