package main

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	left, right := 0, n-1

	for left < n-1 && arr[left] < arr[left+1] {
		left++
	}

	for right > 0 && arr[right] < arr[right-1] {
		right--
	}

	return left == right && left != 0 && right != n-1
}
