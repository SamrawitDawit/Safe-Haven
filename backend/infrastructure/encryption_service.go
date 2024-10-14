package infrastructure

import (
	"backend/domain"
	"backend/utils"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type EncryptionService struct {
	Key string
}

func (e *EncryptionService) Encrypt(plaintext string) (string, *domain.CustomError) {
	// Convert the key into a byte slice
	keyBytes := []byte(e.Key)

	// Generate a new AES cipher block using the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		utils.LogError("Failed to create new AES cipher block", err)
		return "", domain.NewCustomError(err.Error(), 500)
	}

	// Create a byte slice to hold the ciphertext and IV (Initialization Vector)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Generate a random IV and place it at the beginning of the ciphertext slice
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		utils.LogError("Failed to generate random IV", err)
		return "", domain.NewCustomError(err.Error(), 500)
	}

	// Encrypt the plaintext using CFB mode
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// Return the ciphertext as a hexadecimal string
	return hex.EncodeToString(ciphertext), nil
}

// DecryptAES256 decrypts a string using AES-256 decryption
func (e *EncryptionService) Decrypt(ciphertext string) (string, *domain.CustomError) {
	// Convert the key and ciphertext into byte slices
	keyBytes := []byte(e.Key)
	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		utils.LogError("Failed to decode ciphertext", err)
		return "", domain.NewCustomError(err.Error(), 500)
	}

	// Generate a new AES cipher block using the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		utils.LogError("Failed to create new AES cipher block", err)
		return "", domain.NewCustomError(err.Error(), 500)
	}

	// Check if the ciphertext is at least as long as the AES block size
	if len(ciphertextBytes) < aes.BlockSize {
		return "", domain.NewCustomError("ciphertext is too short", 400)
	}

	// Extract the IV from the ciphertext
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	// Decrypt the ciphertext using CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	// Return the plaintext as a string
	return string(ciphertextBytes), nil
}
