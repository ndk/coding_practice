package exercise

func backspaceCompare(s string, t string) bool {
	sb, tb := 0, 0
	si, ti := len(s)-1, len(t)-1
	for {
		if si >= 0 && s[si] == '#' {
			sb++
			si--
			continue
		}
		if ti >= 0 && t[ti] == '#' {
			tb++
			ti--
			continue
		}

		if sb > 0 {
			si--
			sb--
			continue
		}
		if tb > 0 {
			ti--
			tb--
			continue
		}

		if si < 0 || ti < 0 {
			return si < 0 && ti < 0
		}

		if s[si] == t[ti] {
			si--
			ti--
			continue
		}

		return false
	}
}
