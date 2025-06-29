package maps

func canConstruct(ransomNote string, magazine string) bool {
	record := [26]rune{}
	for _, v := range magazine {
		record[v-'a']++
	}

	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}
