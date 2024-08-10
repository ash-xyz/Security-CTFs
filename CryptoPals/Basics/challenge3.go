package main

import (
	"encoding/hex"
	"strings"
)

const challenge3 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

var englishFreq = map[rune]float64{
	'a': 8.167,
	'b': 1.492,
	'c': 2.782,
	'd': 4.253,
	'e': 12.702,
	'f': 2.228,
	'g': 2.015,
	'h': 6.094,
	'i': 6.966,
	'j': 0.153,
	'k': 0.772,
	'l': 4.025,
	'm': 2.406,
	'n': 6.749,
	'o': 7.507,
	'p': 1.929,
	'q': 0.095,
	'r': 5.987,
	's': 6.327,
	't': 9.056,
	'u': 2.758,
	'v': 0.978,
	'w': 2.361,
	'x': 0.150,
	'y': 1.974,
	'z': 0.074,
	' ': 1.0,
}

func scoreText(text string) float64 {
	text = strings.ToLower(text)
	score := 0.0

	for _, char := range text {
		if freq, ok := englishFreq[char]; ok {
			score += freq
		} else {
			score -= 20
		}
	}

	return score
}

func solveChallenge3() string {
	cipherText, err := hex.DecodeString(challenge3)
	if err != nil {
		panic(err)
	}

	possiblePlainText := make([]byte, len(cipherText))
	var plainText string
	plainTextScore := 0.0
	for char := 0; char < 256; char++ {
		for i, r := range cipherText {
			possiblePlainText[i] = r ^ byte(char)
		}
		if plainTextScore < scoreText(string(possiblePlainText)) {
			plainText = string(possiblePlainText)
		}
	}
	return plainText
}
