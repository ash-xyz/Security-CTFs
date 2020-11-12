from pwn import xor, unhex
enc_flag = unhex(
    '73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d')
"""
for i in range(100):
    canidate = xor(enc_flag,i)
    if canidate.decode()[0] == 'c':
        print(canidate)
        print(i)
"""
print(xor(enc_flag, 16))
