package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(n int) string {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
