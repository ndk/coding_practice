package exercise

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MeetingPlanner(slotsA [][]int, slotsB [][]int, dur int) []int {
	for len(slotsA) > 0 && len(slotsB) > 0 {
		begin := max(slotsA[0][0], slotsB[0][0])
		end := min(slotsA[0][1], slotsB[0][1])

		if (end - begin) >= dur {
			return []int{begin, begin + dur}
		}
		if slotsB[0][1] < slotsA[0][1] {
			slotsB = slotsB[1:]
			continue
		}
		slotsA = slotsA[1:]
	}
	return []int{}
}
