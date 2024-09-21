def reverseWords(s):
    return " ".join(s.split()[::-1])


def reverseWords1(s):
    return " ".join(reversed(s.split()))


def reverseWords2(s):
    s = s.strip()
    i = j = len(s) - 1
    res = []
    while i >= 0:
        while i >= 0 and s[i] != " ":
            i -= 1
        res.append(s[i + 1 : j + 1])
        while s[i] == " ":
            i -= 1
        j = i

    return " ".join(res)
