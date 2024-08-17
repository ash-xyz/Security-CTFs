package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	mathRand "math/rand"
)

const CBC = 0
const ECB = 1
const KEY_SIZE = 16

func getRandomByteKey(key_size int) []byte {
	key := make([]byte, key_size)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}

func addPadding(plain []byte, before int, after int) []byte {
	paddingLength := before + after
	paddedBytes := make([]byte, len(plain)+paddingLength)

	for i := 0; i < before; i++ {
		paddedBytes[i] = byte(paddingLength)
	}
	for i := before; i < before+len(plain); i++ {
		paddedBytes[i] = byte(paddingLength)
	}
	for i := len(plain) + before; i < len(paddedBytes); i++ {
		paddedBytes[i] = byte(paddingLength)
	}

	return paddedBytes
}

func ecbEncrypt(block cipher.Block, plainText []byte) []byte {
	blockSize := block.BlockSize()
	encrypted := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], plainText[i:i+blockSize])
	}

	return encrypted
}
func getEncryptedText(plainText []byte) ([]byte, int) {
	key := getRandomByteKey(KEY_SIZE)
	iv := make([]byte, KEY_SIZE)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	before := mathRand.Int()%6 + 5
	after := mathRand.Int()%6 + 5

	paddedText := addPadding(plainText, before, after)
	paddedText = paddingPKCS7(paddedText, KEY_SIZE) // make it blocks of KEY_SIZE

	choice := mathRand.Int() % 2

	switch choice {
	case CBC:
		cbc := cipher.NewCBCEncrypter(block, iv)
		cipherText := make([]byte, len(paddedText))
		cbc.CryptBlocks(cipherText, paddedText)
		return cipherText, CBC
	case ECB:
		return ecbEncrypt(block, paddedText), ECB
	}

	return nil, -1
}

func encryptionOracle(cypherText []byte) int {
	seen := map[string]struct{}{}
	for i := 0; i < len(cypherText); i += KEY_SIZE {
		block := string(cypherText[i : i+KEY_SIZE])
		if _, ok := seen[block]; ok {
			return ECB
		}
		seen[block] = struct{}{}
	}

	return CBC
}

func solveChallenge11() {
	plainText := bytes.Repeat([]byte("There's a certain maddness to this world that's not apparent when doing this as trivial as this"), 5)

	for i := 0; i < 100; i++ {
		cyphertext, encryptionMethod := getEncryptedText(plainText)
		predictedEncryptionMethod := encryptionOracle(cyphertext)
		if encryptionMethod != predictedEncryptionMethod {
			panic("The Oracle is blind!!!")
		}
	}

	fmt.Println("The Oracle sees all!")
}
