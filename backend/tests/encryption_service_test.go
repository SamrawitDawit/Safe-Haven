package tests

import (
	"backend/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EncryptionServiceSuite struct {
	suite.Suite
	service *infrastructure.EncryptionService
}

func (suite *EncryptionServiceSuite) SetupTest() {
	// Example key for AES-256 (32 bytes long)
	suite.service = &infrastructure.EncryptionService{
		Key: "my32characterlongencryptionkey!!",
	}
}

// Test Encrypt

func (suite *EncryptionServiceSuite) TestEncrypt_Success() {
	plaintext := "mySensitiveData"
	ciphertext, err := suite.service.Encrypt(plaintext)

	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), ciphertext)

	// Ciphertext should not be equal to plaintext
	assert.NotEqual(suite.T(), plaintext, ciphertext)
}

func (suite *EncryptionServiceSuite) TestEncrypt_Failure_InvalidKey() {
	// Set a short key (less than 32 bytes for AES-256)
	suite.service.Key = "shortkey"

	plaintext := "mySensitiveData"
	ciphertext, err := suite.service.Encrypt(plaintext)

	// Expecting an error due to invalid key size
	assert.NotNil(suite.T(), err)
	assert.Empty(suite.T(), ciphertext)
}

// Test Decrypt

func (suite *EncryptionServiceSuite) TestDecrypt_Success() {
	plaintext := "mySensitiveData"
	ciphertext, _ := suite.service.Encrypt(plaintext)

	decryptedText, err := suite.service.Decrypt(ciphertext)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), plaintext, decryptedText)
}

func (suite *EncryptionServiceSuite) TestDecrypt_Failure_InvalidCiphertext() {
	// Use an invalid ciphertext (e.g., empty or malformed)
	invalidCiphertext := "invalidciphertext"

	decryptedText, err := suite.service.Decrypt(invalidCiphertext)

	// Expecting an error due to invalid ciphertext
	assert.NotNil(suite.T(), err)
	assert.Empty(suite.T(), decryptedText)
}

// Test Ciphertext too short for Decrypt

func (suite *EncryptionServiceSuite) TestDecrypt_Failure_ShortCiphertext() {
	// Short ciphertext that's smaller than AES block size
	shortCiphertext := "abcd"

	decryptedText, err := suite.service.Decrypt(shortCiphertext)

	// Expecting an error due to short ciphertext
	assert.NotNil(suite.T(), err)
	assert.Empty(suite.T(), decryptedText)
}

func TestEncryptionServiceSuite(t *testing.T) {
	suite.Run(t, new(EncryptionServiceSuite))
}
