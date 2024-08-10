package main

import (
	"bufio"
	"encoding/hex"
	"os"
	"strings"
)

func readHexLinesAsBytes(filePath string) ([][]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hex, err := hex.DecodeString(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, hex)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ScoreText(s string) int {
	score := 0
	goodChars := [12]string{"a", "e", "i", "o", "u", "r", "s", "t", "l", "m", "n", " "}

	for i := range goodChars {
		score += strings.Count(strings.ToLower(s), goodChars[i])
	}

	return score
}

func solveChallenge4() string {
	input, err := readHexLinesAsBytes("input/challenge4.txt")
	if err != nil {
		panic(err)
	}

	var bestPlainText string
	bestScore := 0

	for _, cipherText := range input {
		possiblePlainText := make([]byte, len(cipherText))
		for char := 0; char < 256; char++ {
			for i, r := range cipherText {
				possiblePlainText[i] = r ^ byte(char)
			}
			possibleScore := ScoreText(string(possiblePlainText))
			if possibleScore > bestScore {
				bestPlainText = string(possiblePlainText)
				bestScore = possibleScore
			}
		}
	}
	return bestPlainText
}
