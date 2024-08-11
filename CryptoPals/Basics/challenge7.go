package main

import (
	"crypto/aes"
	"fmt"
)

const challenge7Key = "YELLOW SUBMARINE"

func solveChallenge7() {
	block, err := aes.NewCipher([]byte(challenge7Key))
	if err != nil {
		panic(err)
	}

	cypherText := getBase64ToBytesFromPath("input/7.txt")

	plainText := make([]byte, len(cypherText))
	blockSize := 16
	for i := 0; i <= len(cypherText)-blockSize; i += blockSize {
		block.Decrypt(plainText[i:i+blockSize], cypherText[i:i+blockSize])
	}
	fmt.Println(string(plainText))
}
