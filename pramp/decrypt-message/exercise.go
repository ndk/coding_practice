package exercise

func Decrypt(word string) string {
	if len(word) == 0 {
		return ""
	}

	const alphabetSize = ('z' - 'a') + 1

	result := make([]byte, len(word))
	prev := 1
	for i, c := range word {
		r := int(c) - prev
		m := ('a' - r) / alphabetSize
		if ('a' - r) % alphabetSize > 0 {
			m++
		}
		r += alphabetSize * m
		result[i] = byte(r)
		prev += r
	}
	return string(result)
}
