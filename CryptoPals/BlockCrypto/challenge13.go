package main

import (
	"crypto/aes"
	"fmt"
	"strings"
)

var key = getRandomByteKey(16)

func parseStructuredCookie(cookie string) map[string]string {
	pairs := strings.Split(cookie, "&")
	keyValueMap := make(map[string]string)

	for _, keyValue := range pairs {
		keyValue := strings.Split(keyValue, "=")
		keyValueMap[keyValue[0]] = keyValue[1]
	}
	return keyValueMap
}

func profileFor(email string) string {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")

	return fmt.Sprintf("email=%s&uid=10&role=user", email)
}

func decryptUserProfile(ciphertext []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// Remove padding
	unpadded := unpadPKCS7(plaintext, block.BlockSize())

	return string(unpadded)
}
func padPKCS7(data []byte, blockSize int) []byte {
	if blockSize < 1 || blockSize > 255 {
		panic("invalid blocksize")
	}
	paddingSize := blockSize - (len(data) % blockSize)
	for i := 0; i < paddingSize; i++ {
		data = append(data, byte(paddingSize))
	}
	return data
}

func encryptUserProfile(userProfile string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Pad the input to be a multiple of the block size
	paddedProfile := padPKCS7([]byte(userProfile), aes.BlockSize)

	ciphertext := make([]byte, len(paddedProfile))
	for i := 0; i < len(paddedProfile); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], paddedProfile[i:i+aes.BlockSize])
	}

	return ciphertext
}

func unpadPKCS7(data []byte, blockSize int) []byte {
	if len(data) == 0 || len(data)%blockSize != 0 {
		panic("invalid data length")
	}
	paddingSize := int(data[len(data)-1])
	if paddingSize == 0 || paddingSize > blockSize {
		panic("invalid padding size")
	}
	for i := len(data) - paddingSize; i < len(data); i++ {
		if data[i] != byte(paddingSize) {
			panic("invalid padding")
		}
	}
	return data[:len(data)-paddingSize]
}

func solveChallenge13() {
	payload := "AAAAAAAAAAadmin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b"
	encryptedProfile := encryptUserProfile(profileFor(payload))

	normalProfile := encryptUserProfile(profileFor("pwned@ash.xyz"))

	block0 := normalProfile[0:16]
	block1 := normalProfile[16:32]
	block2 := encryptedProfile[16:32]

	maliciousCiphertext := append(append(block0, block1...), block2...)

	decryptedProfile := decryptUserProfile(maliciousCiphertext)
	fmt.Println(decryptedProfile)

	parsedProfile := parseStructuredCookie(decryptedProfile)
	fmt.Println(parsedProfile)
}
