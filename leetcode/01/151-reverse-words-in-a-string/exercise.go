package exercise

import (
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	b := make([]byte, n)

	offset := 0
	p := 0
	for ; p < n && s[p] == ' '; p++ {
		offset++
	}

	begin, end := p, p
	spaces := 0
	for ; p < n; p++ {
		if s[p] != ' ' {
			end = p
			continue
		}

		spaces = -1
		for ; p < n && s[p] == ' '; p++ {
			spaces++
		}
		if p == n {
			break
		}

		for i, j := begin, end; i <= j; {
			b[i-offset], b[j-offset] = s[j], s[i]
			i++
			j--
		}
		b[end-offset+1] = ' '
		offset += spaces
		begin, end = p, p
	}

	for i, j := begin, end; i <= j; {
		b[i-offset], b[j-offset] = s[j], s[i]
		i++
		j--
	}

	for i, j := 0, end-offset; i <= j; {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b[:end-offset+1])
}

///////////////////////////////////////////////////////////////////////////////
// Approach 1: Built-in Split + Reverse

func reverseWords_split(s string) string {
	s = strings.TrimSpace(s)
	list := strings.Split(s, " ")
	var trimmed []string
	for _, s := range list {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		trimmed = append(trimmed, s)
	}
	for l, r := 0, len(trimmed)-1; l < r; {
		trimmed[l], trimmed[r] = trimmed[r], trimmed[l]
		l++
		r--
	}
	return strings.Join(trimmed, " ")
}

///////////////////////////////////////////////////////////////////////////////
// Approach 2: Reverse the Whole String and Then Reverse Each Word

func reverseWords_double(sa string) string {
	s := []byte(sa)

	reverse := func(b []byte, i, j int) {
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	// reverse the whole string
	reverse(s, 0, len(s)-1)

	n := len(s)
	idx := 0
	for start := 0; start < n; start++ {
		if s[start] != ' ' {
			// go to the beginning of the word
			if idx != 0 {
				s[idx] = ' '
				idx++
			}

			// go to the end of the word
			end := start
			for end < n && s[end] != ' ' {
				s[idx] = s[end]
				idx++
				end++
			}

			// reverse the word
			reverse(s, idx-(end-start), idx-1)

			// move to the next word
			start = end
		}
	}
	return string(s[:idx])
}

///////////////////////////////////////////////////////////////////////////////
// Approach 3: Deque of Words

func reverseWords_deque(s string) string {
	left, right := 0, len(s)-1
	// remove leading spaces
	for left <= right && s[left] == ' ' {
		left++
	}

	// remove trailing spaces
	for left <= right && s[right] == ' ' {
		right--
	}

	var d []string
	var word []byte
	// push word by word in front of deque
	for left <= right {
		c := s[left]

		if len(word) != 0 && c == ' ' {
			d = append([]string{string(word)}, d...)
			word = nil
		} else if c != ' ' {
			word = append(word, c)
		}
		left++
	}
	d = append([]string{string(word)}, d...)

	return strings.Join(d, " ")
}
