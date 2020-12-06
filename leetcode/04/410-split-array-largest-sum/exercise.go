package exercise

func splitArray(nums []int, m int) int {
	lo, hi := nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		hi += nums[i]
		if lo < nums[i] {
			lo = nums[i]
		}
	}
	result := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sum := 0
		pieces := 1
		for _, num := range nums {
			if sum+num > mid {
				pieces++
				if pieces > m {
					lo = mid + 1
					break
				}
				sum = num
			} else {
				sum += num
			}
		}
		if pieces <= m {
			if mid < result {
				result = mid
			}
			hi = mid - 1
		}
	}
	return result
}
