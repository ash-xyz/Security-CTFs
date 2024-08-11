package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

// Dumb solution, some 16 byte instance is going to be the same
// So just loop through and find the hex with the most equal bytes
func solveChallenge8() string {
	file, err := os.Open("input/8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var encryptedBlock string
	blockWithMostRepeats := math.MaxInt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		count := make(map[string]int)

		for i := 0; i < len(line); i += 32 {
			block := line[i : i+32]
			count[block]++
		}

		if len(count) < blockWithMostRepeats {
			blockWithMostRepeats = len(count)
			encryptedBlock = line
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return encryptedBlock
}
