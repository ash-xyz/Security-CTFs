p = 29
candidates = [14, 6, 11]

for candidate in candidates:
    for i in range(29):
        if(pow(i, 2, p) == candidate):
            print(f"Square root: {i}, Quadratic Residue: {candidate}")
