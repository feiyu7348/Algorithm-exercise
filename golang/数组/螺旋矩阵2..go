package main

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1
	matrix := make([][]int, n)
	num := 1
	tar := n * n

	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
