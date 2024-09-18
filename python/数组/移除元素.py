"""
从列表 nums 中移除值为 val 的元素，并返回新的列表长度。

参数：
    nums：一个整数列表。
    val：要移除的整数。

返回值：新的列表长度。

测试用例：
>>> nums = [3,2,2,3]
>>> removeElement(nums, 3)
2
>>> nums
[2, 2]
"""


def removeElement(nums, val):
    # 相反双指针法
    n = len(nums)
    left = 0
    right = n - 1

    while left <= right:
        if nums[left] == val:
            nums[left] = nums[right]
            right -= 1
        else:
            left += 1

    return left


def removeElement1(nums, val):
    # 正向双指针法, 慢指针指向新数组的末尾，快指针遍历整个数组
    n = len(nums)
    slow = 0
    fast = 0

    while fast < n:
        if nums[fast] != val:
            nums[slow] = nums[fast]
            slow += 1

        fast += 1

    return slow


if __name__ == "__main__":
    nums = [3, 2, 2, 3]
    val = 3
    k = removeElement(nums, val)
    print(k)
    print(nums[:k])
