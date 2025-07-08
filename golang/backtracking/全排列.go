package backtracking

func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	var path []int
	used := make(map[int]bool)

	var defs func()
	defs = func() {
		if len(path) == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				defs()
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	defs()
	return res
}
