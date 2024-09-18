def smallerNumbersThanCurrent(nums):
    n = len(nums)
    res = [0] * n
    for i in range(n):
        sum = 0
        for j in range(n):
            if nums[i] > nums[j]:
                sum += 1

        res[i] = sum

    return res


nums = [8, 1, 2, 2, 4]
print(smallerNumbersThanCurrent(nums))
