from pwn import unhex, xor
enc_flag = unhex(
    '0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104')

# Find the partial XOR key
print(f"First part of key:{xor('crypto{',enc_flag[:7])}")
print(f"Last Part of key:{xor('}',enc_flag[-1])}\n")

# Get Full Key
key = 'myXORkey' * int((len(enc_flag)/9))
print(f"Key:{key}\n")

# Get Decrypted Flag
print(f"Flag: {xor(key,enc_flag).decode()}")
