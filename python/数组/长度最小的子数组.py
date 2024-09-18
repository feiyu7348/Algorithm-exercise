def minSubArrayLen(target, nums):
    # 滑动窗口法
    n = len(nums)
    left = 0
    right = 0
    min_len = n + 1
    sum = 0

    while right < n:
        sum += nums[right]

        while sum >= target:
            min_len = min(min_len, right - left + 1)
            sum -= nums[left]
            left += 1

        right += 1

    return min_len if min_len <= n else 0


if __name__ == "__main__":
    target = 7
    nums = [2, 3, 1, 2, 4, 3]
    min_len = minSubArrayLen(target, nums)
    print(min_len)
