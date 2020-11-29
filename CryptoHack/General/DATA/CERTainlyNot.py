from Crypto.PublicKey import RSA
key = RSA.importKey(open("2048b.der", "rb").read())
print(key.n)