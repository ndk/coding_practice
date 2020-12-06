package exercise

func longestPalindrome(s string) string {
	expand := func(s string, left int, right int) int {
		l, r := left, right
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return r - l - 1
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len := max(expand(s, i, i), expand(s, i, i+1))
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}
