package exercise

func isMatch(s string, p string) bool {
	type pair struct {
		i int
		j int
	}
	cache := map[pair]bool{}
	var f func(i, j int) bool
	f = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		if j < 0 {
			return false
		}
		if i < 0 {
			return p[j] == '*' && f(i, j-2)
		}

		if r, ok := cache[pair{i, j}]; ok {
			return r
		}

		if p[j] == '*' {
			c := p[j-1]
			r := f(i, j-2) || ((s[i] == c || c == '.') && f(i-1, j))
			cache[pair{i, j}] = r
			return r
		}

		if p[j] == '.' || p[j] == s[i] {
			r := f(i-1, j-1)
			cache[pair{i, j}] = r
			return r
		}

		return false
	}
	return f(len(s)-1, len(p)-1)
}

///////////////////////////////////////////////////////////////////////////////
// Approach 1: Recursion
func isMatch1(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) != 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return (isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p)))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

///////////////////////////////////////////////////////////////////////////////
// Approach 2: Dynamic Programming

func isMatch2(s string, p string) bool {
	n, m := len(s), len(p)
	memo := make([]*bool, (n+1)*(m+1))
	var dp func(i int, j int, text string, pattern string) bool
	dp = func(i int, j int, s string, p string) bool {
		if memo[i*(m+1)+j] != nil {
			return *memo[i*(m+1)+j] == true
		}
		var ans bool
		if j == len(p) {
			ans = i == len(s)
		} else {
			firstMatch := (i < len(s) && (p[j] == s[i] || p[j] == '.'))

			if j+1 < len(p) && p[j+1] == '*' {
				ans = (dp(i, j+2, s, p) || firstMatch && dp(i+1, j, s, p))
			} else {
				ans = firstMatch && dp(i+1, j+1, s, p)
			}
		}
		if ans {
			t := true
			memo[i*(m+1)+j] = &t
		} else {
			t := false
			memo[i*(m+1)+j] = &t
		}
		return ans
	}
	return dp(0, 0, s, p)
}
