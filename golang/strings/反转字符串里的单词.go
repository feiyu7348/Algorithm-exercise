package MyStrings

func reverseWords(s string) string {
	n := len(s)
	b := []byte(s)
	slow := 0
	for i := 0; i < n; i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < n && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[:slow]

	reverseString(b)
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverseString(b[last:i])
			last = i + 1
		}
	}

	return string(b)
}
