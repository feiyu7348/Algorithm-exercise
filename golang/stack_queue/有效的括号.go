package stackqueue

func IsValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}
	for _, v := range s {
		if _, ok := m[v]; !ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			peek := stack[len(stack)-1]
			if peek != m[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
