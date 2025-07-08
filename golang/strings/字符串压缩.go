package MyStrings

import "strconv"

func compressString(S string) string {
	n := len(S)
	if n == 0 {
		return S
	}

	var newS []byte
	curChar := S[0]
	count := 1
	for i := 1; i < n; i++ {
		if S[i] == curChar {
			count++
		} else {
			newS = append(newS, curChar)
			newS = append(newS, strconv.Itoa(count)...)
			curChar = S[i]
			count = 1
		}
	}

	newS = append(newS, curChar)
	newS = append(newS, strconv.Itoa(count)...)
	newSS := string(newS)

	if len(newSS) >= n {
		return S
	}
	return newSS
}
