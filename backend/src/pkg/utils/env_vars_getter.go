package utils

import (
	"os"
	"strconv"
)

func GetStringEnv(key string, defaultVal string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return defaultVal
}

func GetIntEnv(key string, defaultVal int) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)

	if err == nil {
		return value
	}

	return defaultVal
}
