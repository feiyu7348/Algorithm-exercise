package main

import (
	"fmt"

	"algorithm/sorts"
)

func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	fmt.Println(sorts.Search(nums, 10))
	fmt.Println(sorts.BubbleSort(nums))
	fmt.Println(sorts.SelectSort(nums))
	fmt.Println(sorts.InsertSort(nums))
	fmt.Println(sorts.QuickSort(nums))
}
