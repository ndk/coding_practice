package exercise

import "unicode"

func licenseKeyFormatting(s string, k int) string {
	n := len(s)
	result := make([]byte, n+n/k+1)
	i := n - 1
	p := len(result) - 1
	r := 0
	for ; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if r == k {
			result[p] = '-'
			p--
			r = 0
		}
		c := unicode.ToUpper(rune(s[i]))
		result[p] = byte(c)
		p--
		r++
	}
	return string(result[p+1:])
}
