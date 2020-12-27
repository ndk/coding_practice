package exercise

func missingNumber(nums []int) int {
	diff := int64(0)
	for i, num := range nums {
		diff += int64(num - i)
	}
	return int(int64(len(nums)) - diff)
}

func missingNumber2(nums []int) int {
	sum := int64(0)
	for _, num := range nums {
		sum += int64(num)
	}
	n := int64(len(nums))
	return int(n*(n+1)/2 - sum)
}

///////////////////////////////////////////////////////////////////////////////
// Approach #3 Bit Manipulation

func missingNumber_xor(nums []int) int {
	missing := len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	return missing
}
