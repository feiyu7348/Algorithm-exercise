def rightReverse(s, k):
    s = s[len(s) - k :] + s[0 : len(s) - k]
    return s
