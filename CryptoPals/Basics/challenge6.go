package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strings"
)

func hammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		panic("Strings are of unequal length!")
	}

	dist := 0
	for i := 0; i < len(s1); i++ {
		xor := s1[i] ^ s2[i]
		dist += bits.OnesCount8(uint8(xor))
	}
	return dist
}

func getChallenge6Input() []byte {

	file, err := os.Open("input/6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	base64String := ""
	for scanner.Scan() {
		base64String += strings.TrimSpace(scanner.Text())
	}

	decodedData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return decodedData
}

func transposeBytes(blockSize int, bytes []byte) [][]byte {
	if len(bytes) == 0 || blockSize <= 0 {
		return nil
	}

	numBlocks := (len(bytes) + blockSize - 1) / blockSize
	transposed := make([][]byte, blockSize)

	for i := range transposed {
		transposed[i] = make([]byte, numBlocks)
	}

	for i, b := range bytes {
		row := i % blockSize
		col := i / blockSize
		transposed[row][col] = b
	}

	return transposed
}

func singleCharXOR(block []byte) byte {
	temp := make([]byte, len(block))
	bestScore := 0
	var blockKey byte
	for char := 0; char < 256; char++ {
		for i, r := range block {
			temp[i] = r ^ byte(char)
		}

		curScore := ScoreText(string(temp))
		if bestScore < curScore {
			blockKey = byte(char)
			bestScore = curScore
		}
	}

	return blockKey
}

func solveChallenge6() {
	cypherText := getChallenge6Input()
	keySize := 0
	minScore := math.Inf(1)

	for candidateKeySize := 2; candidateKeySize < 42; candidateKeySize++ {
		prev := cypherText[:candidateKeySize] //overlapping blocks isn't ideal here, but it works so :)
		score := 0.0
		for i := candidateKeySize; i < candidateKeySize*10; i += candidateKeySize {
			cur := cypherText[i : i+candidateKeySize]
			score += float64(hammingDistance(string(prev), string(cur))) / float64(candidateKeySize)
			prev = cur
		}
		if score < minScore {
			minScore = score
			keySize = candidateKeySize
		}
	}

	transposition := transposeBytes(keySize, cypherText)

	key := []byte{}
	for i := 0; i < len(transposition); i++ {
		key = append(key, singleCharXOR(transposition[i]))
	}
	fmt.Println(string(repeatingKeyXOR(cypherText, key)))
}
