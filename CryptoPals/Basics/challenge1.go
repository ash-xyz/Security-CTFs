package main

import (
	"encoding/base64"
	"encoding/hex"
)

// 1. Convert hex to base64

const challenge1 = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func solveChallenge1() string {
	rawBytes, err := hex.DecodeString(challenge1)
	if err != nil {
		panic(err)
	}

	base64 := base64.RawStdEncoding.EncodeToString(rawBytes)
	return base64
}
