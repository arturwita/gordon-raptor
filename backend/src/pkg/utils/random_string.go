package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(n int) string {
	if n <= 0 {
		return ""
	}

	bytes := make([]byte, (n + 1) / 2)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)[:n]
}
