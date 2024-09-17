def search(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] > target:
            right = mid - 1
        else:
            left = mid + 1

    return -1


def bubble_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        flag = False
        for j in range(0, n - i - 1):
            if nums[j] > nums[j + 1]:
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
                flag = True

        if not flag:
            break

    return nums


def select_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        min_index = i
        for j in range(i + 1, n):
            if nums[j] < nums[min_index]:
                min_index = j

        nums[i], nums[min_index] = nums[min_index], nums[i]

    return nums


def insert_sort(nums):
    n = len(nums)
    for i in range(1, n):
        base = nums[i]
        j = i - 1
        while j >= 0 and nums[j] > base:
            nums[j + 1] = nums[j]
            j -= 1

        nums[j + 1] = base

    return nums


def quick_sort(nums):
    n = len(nums)
    if n < 2:
        return nums

    pivot = nums[0]
    left = [nums[i] for i in range(1, n) if nums[i] < pivot]
    right = [nums[i] for i in range(1, n) if nums[i] >= pivot]

    return quick_sort(left) + [pivot] + quick_sort(right)


if __name__ == "__main__":
    nums = [1, 3, 4, 5, 2, 7, 6]
    print(quick_sort(nums))
