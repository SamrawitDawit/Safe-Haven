package tests

import (
	"backend/domain"
	"backend/infrastructure"
	"backend/usecases/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PasswordServiceSuite struct {
	suite.Suite
	service interfaces.HashingServiceInterface
}

func (suite *PasswordServiceSuite) SetupTest() {
	suite.service = &infrastructure.HashingService{}
}

// Test HashPassword

func (suite *PasswordServiceSuite) TestHashPassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, err := suite.service.HashPassword(password)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), hashedPassword)

	// Verify that the hashed password is different from the plain password
	assert.NotEqual(suite.T(), password, hashedPassword)
}

// Test ComparePassword

func (suite *PasswordServiceSuite) TestComparePassword_Success() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	err := suite.service.CheckPasswordHash(password, hashedPassword)
	suite.Nil(err)
}

func (suite *PasswordServiceSuite) TestComparePassword_Failure() {
	password := "mySecureP@ssw0rd"
	hashedPassword, _ := suite.service.HashPassword(password)

	// Incorrect password should not match
	err := suite.service.CheckPasswordHash("wrongPassword", hashedPassword)
	suite.Equal(err, domain.ErrInvalidPassword)
}

func TestPasswordServiceSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceSuite))
}
