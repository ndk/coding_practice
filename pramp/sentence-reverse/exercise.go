package exercise

func reverseSentence(sentence string) string {
	words := []int{0}
	for i, c := range sentence {
		if c == ' ' {
			words = append(words, i+1)
		}
	}

	result := make([]byte, len(sentence)+1)
	p := 0
	for w := len(words) - 1; w >= 0; w-- {
		for i := words[w]; i < len(sentence) && sentence[i] != ' '; i++ {
			result[p] = sentence[i]
			p++
		}
		result[p] = ' '
		p++
	}

	return string(result[:len(result)-1])
}
