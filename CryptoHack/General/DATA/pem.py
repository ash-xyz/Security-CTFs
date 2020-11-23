#https://pycryptodome.readthedocs.io/en/latest/src/public_key/rsa.html
from Crypto.PublicKey import RSA
key = RSA.importKey(open("pem.pem", "rb").read())
print(key.d)