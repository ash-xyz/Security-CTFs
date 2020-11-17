def extended_gcd(p, q):
    if (p == 0):
        return 0, 1

    x, y = extended_gcd(q % p, p)

    x2 = y - (q//p) * x
    y2 = x

    return x2, y2


print(extended_gcd(26513, 32321))
