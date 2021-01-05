package exercise

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	result := 0
	for p1, p2 := 0, 1; p2 < len(nums); {
		if p1 == p2 {
			p2++
			continue
		}
		diff := nums[p2] - nums[p1]
		if diff == k {
			result++
			for p := nums[p1]; p1 < len(nums) && p == nums[p1]; p1++ {
			}
			for p := nums[p2]; p2 < len(nums) && p == nums[p2]; p2++ {
			}
		} else if diff < k {
			p2++
		} else {
			p1++
		}
	}

	return result
}

func findPairs2(nums []int, k int) int {
	result := 0

	counter := map[int]int{}
	for _, n := range nums {
		counter[n]++
	}

	for x, val := range counter {
		if k > 0 {
			if _, ok := counter[x+k]; ok {
				result++
			}
		} else if k == 0 && val > 1 {
			result++
		}
	}
	return result
}
