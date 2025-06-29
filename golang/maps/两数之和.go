package maps

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		} else {
			m[v] = i
		}
	}
	return nil
}
