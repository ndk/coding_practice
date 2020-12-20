package exercise

func reverseWords(s []byte) {
	n := len(s)

	begin := 0
	for end := 1; end < n; end++ {
		if s[end] == ' ' {
			for l, r := begin, end-1; l <= r; {
				s[l], s[r] = s[r], s[l]
				l++
				r--
			}
			begin = end + 1
		}
	}

	n--
	for l, r := begin, n; l <= r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	for l := 0; l <= n; {
		s[l], s[n] = s[n], s[l]
		l++
		n--
	}
}
