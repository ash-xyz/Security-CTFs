package main

import "encoding/hex"

// 2. Fixed XOR

func solveChallenge2() string {
	x1, x2 := "1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"
	rawX1, err := hex.DecodeString(x1)
	if err != nil {
		panic(err)
	}
	rawX2, err := hex.DecodeString(x2)
	if err != nil {
		panic(err)
	}

	xor := []byte{}
	for i := 0; i < len(rawX1); i++ {
		xor = append(xor, rawX1[i]^rawX2[i])
	}

	return hex.EncodeToString(xor)
}
