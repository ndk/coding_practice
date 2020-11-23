package exercise

func IndexEqualsValueSearch(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r - l) / 2
		switch {
		case arr[m] < m:
			l = m + 1
		case arr[m] > m:
			r = m - 1
		default:
			r = m
		}
	}
	if arr[l] == l {
		return l
	}
	return -1
}

func IndexEqualsValueSearch_linear(arr []int) int {
	for i := 0; i < len(arr) && arr[i] <= i; i++ {
		if arr[i] == i {
			return i
		}
	}
	return -1
}
