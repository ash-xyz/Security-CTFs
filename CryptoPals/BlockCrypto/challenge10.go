package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const challenge10Key = "YELLOW SUBMARINE"

func getBytesFromPath(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data += strings.TrimSpace(scanner.Text())
	}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return []byte(decoded)
}

func xor(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i%len(key)]
	}
	return ciphertext
}

func solveChallenge10() {
	aes, err := aes.NewCipher([]byte(challenge10Key))
	if err != nil {
		panic(err)
	}
	c_i := make([]byte, 16) // IV

	for i := 0; i < 16; i++ {
		c_i[i] = 0
	}

	cypherText := getBytesFromPath("input/10.txt")
	plainText := make([]byte, len(cypherText))

	for i := 0; i < len(cypherText); i += 16 {
		d_k_c_i := make([]byte, 16)
		aes.Decrypt(d_k_c_i, cypherText[i:i+16])
		copy(plainText[i:i+16], xor(d_k_c_i, c_i))
		c_i = cypherText[i : i+16]
	}

	fmt.Println(string(plainText))
}
