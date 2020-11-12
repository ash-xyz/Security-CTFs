from pwn import *  # pip install pwntools
import json
import base64
import codecs
from Crypto.Util.number import bytes_to_long, long_to_bytes

r = remote('socket.cryptohack.org', 13377, level='debug')


def json_recv():
    line = r.recvline()
    return json.loads(line.decode())


def json_send(hsh):
    request = json.dumps(hsh).encode()
    r.sendline(request)


def decode_base64(enc):
    return base64.b64decode(enc).decode()


def decode_hex(enc):
    return unhex(enc).decode()


def decode_rot13(enc):
    return codecs.decode(enc, 'rot_13')


def decode_bigint(enc):
    return long_to_bytes(int(enc, 16)).decode()


def decode_utf(enc):
    return "".join(chr(o) for o in enc)


for i in range(100):
    received = json_recv()

    print("Received type: ")
    print(received["type"])
    print("Received encoded value: ")
    print(received["encoded"])
    message = ""
    if received["type"] == 'base64':
        message = decode_base64(received["encoded"])
    if received["type"] == 'hex':
        message = decode_hex(received["encoded"])
    if received["type"] == 'rot13':
        message = decode_rot13(received["encoded"])
    if received["type"] == 'bigint':
        message = decode_bigint(received["encoded"])
    if received["type"] == 'utf-8':
        message = decode_utf(received["encoded"])
    print(f"Message: {message}")
    to_send = {
        "decoded": message
    }
    json_send(to_send)

json_recv()
