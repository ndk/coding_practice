package exercise

func decodeString(s string) string {
	var i int

	var decode func() []byte
	decode = func() []byte {
		var result []byte

		for ; i < len(s); i++ {
			if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
				result = append(result, s[i])
				continue
			}
			if s[i] == ']' {
				return result
			}
			amount := 0
			for ; '0' <= s[i] && s[i] <= '9'; i++ {
				amount = amount*10 + int(s[i]-'0')
			}
			i++ // '[' is assumed
			subresult := decode()
			for ; amount > 0; amount-- {
				result = append(result, subresult...)
			}
		}

		return result
	}

	return string(decode())
}
