package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func bubble_sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return nums
}

func select_sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}

		nums[i], nums[min_index] = nums[min_index], nums[i]
	}

	return nums
}

func insert_sort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		base := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > base; j-- {
			nums[j+1] = nums[j]
		}

		nums[j+1] = base
	}

	return nums
}

func quick_sort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	pivot := nums[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < n; i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	return append(quick_sort(left), append([]int{pivot}, quick_sort(right)...)...)
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	fmt.Println(search(nums, 10))
	fmt.Println(bubble_sort(nums))
	fmt.Println(select_sort(nums))
	fmt.Println(insert_sort(nums))
	fmt.Println(quick_sort(nums))
}
