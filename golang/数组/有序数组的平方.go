package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	i, j, k := 0, n-1, n-1
	for i <= j {
		left, right := nums[i]*nums[i], nums[j]*nums[j]
		if left > right {
			result[k] = left
			i++
		} else {
			result[k] = right
			j--
		}
		k--
	}

	return result
}
