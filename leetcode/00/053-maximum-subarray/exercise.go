package exercise

func maxSubArray(nums []int) int {
	maxSum, currentSum := nums[0], 0
	for _, num := range nums {
		currentSum += num
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}
