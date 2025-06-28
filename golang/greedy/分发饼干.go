package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	count := 0
	index := n - 1
	for i := m - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			count++
			index--
		}
	}
	return count
}
