def generateMartrix(n):
    res = [[0] * n for _ in range(n)]
    left, right, top, bottom = 0, n - 1, 0, n - 1
    num = 1

    while num <= n * n:
        for i in range(left, right + 1):
            res[top][i] = num
            num += 1
        top += 1

        for i in range(top, bottom + 1):
            res[i][right] = num
            num += 1
        right -= 1

        for i in range(right, left - 1, -1):
            res[bottom][i] = num
            num += 1
        bottom -= 1

        for i in range(bottom, top - 1, -1):
            res[i][left] = num
            num += 1
        left += 1

    return res


n = 3
print(generateMartrix(n))
