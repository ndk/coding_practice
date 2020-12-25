package exercise

func isMatch(text string, pattern string) bool {
	type pair struct {
		t int
		p int
	}
	cache := map[pair]bool{}
	var f func(t, p int) bool
	f = func(t, p int) bool {
		if t < 0 && p < 0 {
			return true
		}
		if t < 0 || p < 0 {
			return false
		}
		if r, ok := cache[pair{t, p}]; ok {
			return r
		}

		if pattern[p] == '*' {
			c := pattern[p-1]
			for ; t < len(text) && text[t] == c; t-- {
				if f(t, p-2) {
					cache[pair{t, p}] = true
					return true
				}
			}
			if f(t, p-2) {
				cache[pair{t, p}] = true
				return true
			}
			return false
		}

		if pattern[p] == '.' || pattern[p] == text[t] {
			r := f(t-1, p-1)
			cache[pair{t, p}] = r
			return r
		}

		return false
	}
	return f(len(text)-1, len(pattern)-1)
}
