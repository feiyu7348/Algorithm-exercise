package backtracking

func CombinetionSum3(k, n int) [][]int {
	var res [][]int
	var path []int
	var defs func(int, int)
	defs = func(start, sum int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n || 9-i+1 < k-len(path) {
				break
			}

			path = append(path, i)
			defs(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	defs(1, 0)
	return res
}
