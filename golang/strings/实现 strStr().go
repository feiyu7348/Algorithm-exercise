package MyStrings

func strStr(haystack string, needle string) int {
	n := len(needle)
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 1; i < n; i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}

		if needle[i] == needle[j+1] {
			j++
		}
		next[i] = j
	}

	j = -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == n-1 {
			return i - n + 1
		}
	}

	return -1
}
