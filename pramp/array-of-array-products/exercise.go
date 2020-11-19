package exercise

func ArrayOfArrayProducts(arr []int) []int {
	if len(arr) < 2 {
		return []int{}
	}

	result := make([]int, len(arr))
	{
		product := 1
		for i := 0; i < len(result); i++ {
			result[i] = product
			product *= arr[i]
		}
	}
	{
		product := 1
		for i := len(arr) - 1; i >= 0; i-- {
			result[i] *= product
			product *= arr[i]
		}
	}
	return result
}
