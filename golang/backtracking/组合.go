package backtracking

func Combine(n int, k int) [][]int {
	var res [][]int
	var path []int

	var defs func(int)
	defs = func(start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}

			path = append(path, i)
			defs(i + 1)
			path = path[:len(path)-1]
		}
	}

	defs(1)
	return res
}
