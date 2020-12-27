package exercise

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for 0 < nums[i] && nums[i] < n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	i := 0
	for ; i < n && i == nums[i]-1; i++ {
	}
	return i + 1
}
