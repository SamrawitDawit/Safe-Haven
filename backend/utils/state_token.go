package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateStateToken generates a random string to use as the state token
func GenerateStateToken() (string, error) {
	b := make([]byte, 32) // 32 bytes generate a 256-bit token
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
