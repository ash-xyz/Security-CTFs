package main

import (
	"encoding/hex"
	"fmt"
)

var challenge5 = [][]byte{
	[]byte("Burning 'em, if you ain't quick and nimble"),
	[]byte("I go crazy when I hear a cymbal"),
}

func repeatingKeyXOR(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i%len(key)]
	}
	return ciphertext
}

func solveChallenge5() {
	key := []byte("ICE")

	for _, line := range challenge5 {
		fmt.Println(hex.EncodeToString(repeatingKeyXOR(line, key)))
	}
}
