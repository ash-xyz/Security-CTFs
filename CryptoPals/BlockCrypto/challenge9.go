package main

func paddingPKCS7(plain []byte, blockLength int) []byte {
	paddingLength := blockLength - (len(plain) % blockLength)
	if paddingLength == 0 {
		return plain
	}
	paddedBytes := make([]byte, len(plain)+paddingLength)
	copy(paddedBytes, plain)

	for i := len(plain); i < len(paddedBytes); i++ {
		paddedBytes[i] = byte(paddingLength)
	}

	return paddedBytes
}
