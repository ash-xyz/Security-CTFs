from pwn import unhex, xor
key1 = unhex('a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313')
key2Xkey3 = unhex('c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1')
enc_flag = unhex('04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf')

print(xor(enc_flag, key1, key2Xkey3))
