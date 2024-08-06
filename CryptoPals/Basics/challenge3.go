package main

import (
	"encoding/hex"
)

const challenge3 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func solveChallenge3() string {
	cipherText, err := hex.DecodeString(challenge3)
	if err != nil {
		panic(err)
	}

	plainText := make([]byte, len(cipherText))

	for i, r := range cipherText {
		plainText[i] = r ^ byte('X') // I'm just being lazy
	}

	return string(plainText)
}
