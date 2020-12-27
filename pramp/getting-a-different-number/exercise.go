package exercise

func getDifferentNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for 0 <= nums[i] && nums[i] < n && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	i := 0
	for ; i < n && i == nums[i]; i++ {
	}
	return i
}
