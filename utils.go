package main

import (
	"crypto/sha256"
	"fmt"
)

func hashString(data string) string {
	hash := sha256.New().Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}
