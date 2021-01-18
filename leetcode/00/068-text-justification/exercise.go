package exercise

import (
	"bytes"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string

	buf := make([]byte, maxWidth)
	for p := 0; p < len(words); {
		prev := p

		width := 0
		for ; p < len(words) && width+len(words[p])+(p-prev) <= maxWidth; p++ {
			width += len(words[p])
		}

		i := 0
		for j := 0; j < len(words[prev]); j++ {
			buf[i] = words[prev][j]
			i++
		}
		if (p - prev) > 1 {
			diff := maxWidth - width
			gap, extraGap := diff/(p-prev-1), diff%(p-prev-1)
			if p == len(words) {
				gap = 1
			}
			for prev++; prev < p; prev++ {
				for to := i + gap; i < to; i++ {
					buf[i] = ' '
				}
				if p < len(words) && extraGap > 0 {
					buf[i] = ' '
					extraGap--
					i++
				}
				for j := 0; j < len(words[prev]); j++ {
					buf[i] = words[prev][j]
					i++
				}
			}
		}
		for ; i < len(buf); i++ {
			buf[i] = ' '
		}

		result = append(result, string(buf))
	}

	return result
}

func fullJustify2(words []string, maxWidth int) []string {
	var result []string

	buf := make([]byte, maxWidth)
	for p := 0; p < len(words); {
		prev := p

		width := 0
		for ; p < len(words) && width+len(words[p])+(p-prev) <= maxWidth; p++ {
			width += len(words[p])
		}

		i := 0
		for j := 0; j < len(words[prev]); j++ {
			buf[i] = words[prev][j]
			i++
		}
		if (p - prev) > 1 {
			extraGap := (maxWidth - width) % (p - prev - 1)
			gap := (maxWidth - width) / (p - prev - 1)
			if p == len(words) {
				gap = 1
			}
			for prev++; prev < p; prev++ {
				for to := i + gap; i < to; i++ {
					buf[i] = ' '
				}
				if p < len(words) {
					if extraGap > 0 {
						buf[i] = ' '
						extraGap--
						i++
					}
				}
				for j := 0; j < len(words[prev]); j++ {
					buf[i] = words[prev][j]
					i++
				}
			}
		}
		for ; i < len(buf); i++ {
			buf[i] = ' '
		}

		result = append(result, string(buf))
	}

	return result
}

func fullJustify3(words []string, maxWidth int) []string {
	var result []string
	{
		lines := 0
		for p := 0; p < len(words); {
			prev := p

			width := 0
			for ; p < len(words) && width+len(words[p])+(p-prev) <= maxWidth; p++ {
				width += len(words[p])
			}

			lines++
		}
		result = make([]string, lines)
	}
	r := 0

	buf := make([]byte, maxWidth)
	for p := 0; p < len(words); {
		prev := p

		width := 0
		for ; p < len(words) && width+len(words[p])+(p-prev) <= maxWidth; p++ {
			width += len(words[p])
		}

		i := 0
		for j := 0; j < len(words[prev]); j++ {
			buf[i] = words[prev][j]
			i++
		}
		if (p - prev) > 1 {
			g := (maxWidth - width) % (p - prev - 1)
			gap := (maxWidth - width) / (p - prev - 1)
			if p == len(words) {
				gap = 1
			}
			for prev++; prev < p; prev++ {
				for to := i + gap; i < to; i++ {
					buf[i] = ' '
				}
				if p < len(words) {
					if g > 0 {
						buf[i] = ' '
						g--
						i++
					}
				}
				for j := 0; j < len(words[prev]); j++ {
					buf[i] = words[prev][j]
					i++
				}
			}
		}
		for ; i < len(buf); i++ {
			buf[i] = ' '
		}

		result[r] = string(buf)
		r++
	}

	return result
}

func fullJustify_leet(words []string, maxWidth int) []string {
	justify := []string{}
	current, curLength := []string{}, 0
	for i, w := range words {
		if curLength+len(current)+len(w) > maxWidth {
			if len(current) == 1 {
				// only one word, all spaces are to the right of the word
				curLine := current[0] + strings.Repeat(" ", maxWidth-len(current[0]))
				justify = append(justify, curLine)
			} else {
				diff := maxWidth - curLength
				spaces := diff / (len(current) - 1)
				more := diff % (len(current) - 1)
				curLine := bytes.Buffer{}
				for ci, cw := range current {
					curLine.WriteString(cw)
					if ci != len(current)-1 {
						moreBlanks := 0
						if more > 0 {
							moreBlanks = 1
							more--
						}
						curLine.WriteString(strings.Repeat(" ", spaces+moreBlanks))
					}
				}
				justify = append(justify, curLine.String())
			}
			current, curLength = []string{}, 0
		}

		curLength += len(w)
		current = append(current, w)

		// last line, left justified and no extra space is inserted between words
		if i == len(words)-1 {
			lastLine := strings.Join(current, " ")
			lastLine = lastLine + strings.Repeat(" ", maxWidth-len(lastLine))
			justify = append(justify, lastLine)
		}
	}
	return justify
}
