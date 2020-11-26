package exercise

func lengthOfLongestSubstring(s string) int {
	last := [128]int{}
	longest, start := 0, 0
	for i, c := range s {
		if last[c] > start {
			start = last[c]
		} else if d := i - start + 1; d > longest {
			longest = d
		}
		last[c] = i + 1
	}
	return longest
}
