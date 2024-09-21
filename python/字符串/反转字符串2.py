def reverseString1(s, k):
    n = len(s)

    res = list(s)
    for i in range(0, n, 2 * k):
        res[i : i + k] = res[i : i + k][::-1]

    return "".join(res)


def reverseString2(s, k):
    n = len(s)

    res = list(s)
    for i in range(0, n, 2 * k):
        left, right = i, min(i + k - 1, n - 1)
        while left < right:
            res[left], res[right] = s[right], s[left]
            left += 1
            right -= 1
    return "".join(res)
