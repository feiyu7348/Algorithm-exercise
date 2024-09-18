def validMountainArray(arr):
    if len(arr) < 3:
        return False

    i = 1
    while i < len(arr) and arr[i] > arr[i - 1]:
        i += 1
    if i == 1 or i == len(arr):
        return False
    while i < len(arr) and arr[i] < arr[i - 1]:
        i += 1

    return i == len(arr)


def validMountainArray2(arr):
    if len(arr) < 3:
        return False

    left, right = 0, len(arr) - 1

    while left < len(arr) - 1 and arr[left] < arr[left + 1]:
        left += 1

    while right > 0 and arr[right] < arr[right - 1]:
        right -= 1

    return left == right and left != 0 and right != len(arr) - 1


arr = [0, 3, 2, 1]
print(validMountainArray(arr))
