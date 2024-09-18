def sortedSquares(nums):
    # 反向双指针法
    n = len(nums)
    left = 0
    right = n - 1
    res = [0] * n

    while left <= right:
        if nums[left] ** 2 > nums[right] ** 2:
            res[n - 1] = nums[left] ** 2
            left += 1
        else:
            res[n - 1] = nums[right] ** 2
            right -= 1

        n -= 1

    return res


nums = [-4, -1, 0, 3, 10]
print(sortedSquares(nums))
