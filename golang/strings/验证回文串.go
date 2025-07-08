package MyStrings

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	l, r := 0, n-1
	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaNum(s byte) bool {
	if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
		return true
	}
	return false
}
