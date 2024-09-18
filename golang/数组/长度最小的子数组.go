package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, right := 0, 0
	sum := 0
	minLen := n + 1

	for right < n {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}

			sum -= nums[left]
			left++
		}

		right++
	}

	if minLen == n+1 {
		return 0
	}

	return minLen
}
