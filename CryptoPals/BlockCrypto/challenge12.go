package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

var (
	unknownKey    []byte
	unknownString []byte
)

func eO(plaintext []byte) []byte {
	input := append(plaintext, unknownString...)
	paddedInput := paddingPKCS7(input, 16)

	block, _ := aes.NewCipher(unknownKey)
	ciphertext := ecbEncrypt(block, paddedInput)

	return ciphertext
}

func detectBlockSize() int {
	initialLength := len(eO([]byte{}))
	for i := 1; ; i++ {
		input := bytes.Repeat([]byte("A"), i)
		ciphertext := eO(input)
		if len(ciphertext) > initialLength {
			return len(ciphertext) - initialLength
		}
	}
}

func detectECB(blockSize int) bool {
	input := bytes.Repeat([]byte("A"), 3*blockSize)
	ciphertext := eO(input)
	return bytes.Equal(ciphertext[blockSize:2*blockSize], ciphertext[2*blockSize:3*blockSize])
}

func decryptByte(known []byte, blockSize int) byte {
	knownLen := len(known)
	padding := bytes.Repeat([]byte("A"), blockSize-1-(knownLen%blockSize))
	target := eO(padding)[:blockSize+knownLen/blockSize*blockSize]

	for i := 0; i < 256; i++ {
		guess := append(padding, append(known, byte(i))...)
		encrypted := eO(guess)[:blockSize+knownLen/blockSize*blockSize]
		if bytes.Equal(encrypted, target) {
			return byte(i)
		}
	}
	return 0
}

func solveChallenge12() []byte {
	unknownKey = getRandomByteKey(16)
	unknownStringBase64 := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	unknownString, _ = base64.StdEncoding.DecodeString(unknownStringBase64)

	blockSize := detectBlockSize()
	fmt.Printf("Detected block size: %d\n", blockSize)

	isECB := detectECB(blockSize)
	fmt.Printf("ECB mode detected: %v\n", isECB)

	var decrypted []byte
	for i := 0; i < len(unknownString); i++ {
		decrypted = append(decrypted, decryptByte(decrypted, blockSize))
	}

	return decrypted
}
