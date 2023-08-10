package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func hashString(data string) string {
	hash := sha256.New().Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func env(key string, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	} else {
		return defaultVal
	}
}
