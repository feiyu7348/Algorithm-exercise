package MyStrings

func reverseStr(s string, k int) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n; i = i + 2*k {
		if i+k <= n {
			reverseString(b[i : i+k])
		} else {
			reverseString(b[i:n])
		}
	}

	return string(b)
}
