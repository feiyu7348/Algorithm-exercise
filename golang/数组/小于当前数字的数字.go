package main

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if nums[i] > nums[j] {
				sum++
			}
		}

		res = append(res, sum)
	}

	return res
}
