package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes   = "0123456789"
	specialBytes  = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
	allBytes      = letterBytes + numberBytes + specialBytes
)

// GenerateComplexRandomString generates a random string of specified length with a mix of letters, numbers, and special characters.
func GenerateComplexRandomString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(allBytes))))
		if err != nil {
			return "", err
		}
		result[i] = allBytes[index.Int64()]
	}
	return string(result), nil
}

