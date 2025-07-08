package MyStrings

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}

	if next[n-1] != -1 && n%(n-next[n-1]-1) == 0 {
		return true
	}
	return false
}
