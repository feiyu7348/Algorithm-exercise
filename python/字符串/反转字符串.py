def reverseStr(s):
    left, right = 0, len(s) - 1
    while left < right:
        s[left], s[right] = s[right], s[left]
        left += 1
        right -= 1

    return s


def reverseStr1(s):
    stack = []
    for i in s:
        stack.append(i)
    s = []
    while stack:
        s.append(stack.pop())
    return s


def reverseStr2(s):
    return s[::-1]


def reverseStr3(s):
    return "".join(reversed(s))
