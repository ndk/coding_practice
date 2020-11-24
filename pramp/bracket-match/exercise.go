package exercise

func BracketMatch(text string) int {
	opened, closed := 0, 0
	for _, c := range text {
		switch c {
		case '(':
			opened++
		case ')':
			if opened > 0 {
				opened--
			} else {
				closed++
			}
		}
	}
	return closed + opened
}
