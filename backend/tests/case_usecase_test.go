package tests

import (
	"backend/domain"
	"backend/tests/mocks"
	"backend/usecases"
	"backend/usecases/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CaseUseCaseTestSuite struct {
	suite.Suite
	caseUseCase        usecases.CaseUseCaseInterface
	mockCaseRepo       *mocks.CaseRepositoryInterface
	mockEncryptService *mocks.EncryptionServiceInterface
}

func (suite *CaseUseCaseTestSuite) SetupTest() {
	suite.mockCaseRepo = new(mocks.CaseRepositoryInterface)
	suite.mockEncryptService = new(mocks.EncryptionServiceInterface)
	suite.caseUseCase = usecases.NewCaseUseCase(
		suite.mockCaseRepo,
		suite.mockEncryptService,
	)
}

func (suite *CaseUseCaseTestSuite) TearDownTest() {
	suite.mockCaseRepo.AssertExpectations(suite.T())
	suite.mockEncryptService.AssertExpectations(suite.T())
}

func (suite *CaseUseCaseTestSuite) TestValidateCaseDto_Success() {
	// Arrange
	caseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	// Act
	err := suite.caseUseCase.ValidateCaseDto(caseDto)

	// Assert
	suite.Nil(err)
}

func (suite *CaseUseCaseTestSuite) TestValidateCaseDto_Fail() {
	// Arrange
	caseDto := dto.CaseDto{
		Title:       "",
		Description: "",
	}

	// Act
	err := suite.caseUseCase.ValidateCaseDto(caseDto)

	// Assert
	suite.Equal(domain.ErrIncompleteCaseInformation, err)
}

func (suite *CaseUseCaseTestSuite) TestDecryptField_Success() {
	// Arrange
	field := "encrypted_field"

	// Expect encryption service to decrypt field successfully
	suite.mockEncryptService.On("Decrypt", field).Return("decrypted_field", nil)

	// Act
	decryptedField, err := suite.caseUseCase.DecryptField(field)

	// Assert
	suite.Nil(err)
	suite.Equal("decrypted_field", decryptedField)
}

func (suite *CaseUseCaseTestSuite) TestDecryptField_Fail() {
	// Arrange
	field := "encrypted_field"

	// Expect encryption service to return an error
	suite.mockEncryptService.On("Decrypt", field).Return("", domain.ErrDecryptionFailed)

	// Act
	decryptedField, err := suite.caseUseCase.DecryptField(field)

	// Assert
	suite.Equal(domain.ErrDecryptionFailed, err)
	suite.Empty(decryptedField)
}

func (suite *CaseUseCaseTestSuite) TestDecrypt_Success() {
	// Arrange
	cases := []*domain.Case{
		{
			Title:       "encrypted_title",
			Description: "encrypted_description",
			ImageURL:    "encrypted_image_url",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil).Times(3)

	// Act
	decryptedCases, err := suite.caseUseCase.Decrypt(cases)

	// Assert
	suite.Nil(err)
	suite.Len(decryptedCases, 1)
	suite.Equal("decrypted_field", decryptedCases[0].Title)
	suite.Equal("decrypted_field", decryptedCases[0].Description)
	suite.Equal("decrypted_field", decryptedCases[0].ImageURL)
}

func (suite *CaseUseCaseTestSuite) TestDecrypt_Fail() {
	// Arrange
	cases := []*domain.Case{
		{
			Title:       "encrypted_title",
			Description: "encrypted_description",
			ImageURL:    "encrypted_image_url",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("", domain.ErrDecryptionFailed)

	// Act
	decryptedCases, err := suite.caseUseCase.Decrypt(cases)

	// Assert
	suite.Equal(domain.ErrDecryptionFailed, err)
	suite.Empty(decryptedCases)
}
func (suite *CaseUseCaseTestSuite) TestCreateCase_Success() {
	// Arrange
	caseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	// Expect encryption service to encrypt all fields successfully
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted_field", nil).Times(3)

	// Expect case repo to create case successfully
	suite.mockCaseRepo.On("CreateCase", mock.Anything).Return(nil)

	// Act
	_, err := suite.caseUseCase.CreateCase(caseDto)

	// Assert
	suite.Nil(err)
}

func (suite *CaseUseCaseTestSuite) TestCreateCase_Fail() {
	// Arrange
	caseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	// Expect encryption service to encrypt all fields successfully
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted_field", nil).Times(3)

	// Expect case repo to return an error
	suite.mockCaseRepo.On("CreateCase", mock.Anything).Return(domain.ErrCaseCreationFailed)

	// Act
	_, err := suite.caseUseCase.CreateCase(caseDto)

	// Assert
	suite.Equal(domain.ErrCaseCreationFailed, err)
}

func (suite *CaseUseCaseTestSuite) TestDeleteCase_Success() {
	// Arrange
	caseID := uuid.New()

	// Expect case repo to delete case successfully
	suite.mockCaseRepo.On("DeleteCase", caseID).Return(nil)

	// Act
	err := suite.caseUseCase.DeleteCase(caseID)

	// Assert
	suite.Nil(err)
}

func (suite *CaseUseCaseTestSuite) TestDeleteCase_Fail() {
	// Arrange
	caseID := uuid.New()

	// Expect case repo to return an error
	suite.mockCaseRepo.On("DeleteCase", caseID).Return(domain.ErrCaseDeletionFailed)

	// Act
	err := suite.caseUseCase.DeleteCase(caseID)

	// Assert
	suite.Equal(domain.ErrCaseDeletionFailed, err)
}

func (suite *CaseUseCaseTestSuite) TestGetAllCases_Success() {
	// Arrange
	cases := []*domain.Case{
		{
			Title: "Test Case",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil)

	// Expect case repo to return all cases successfully
	suite.mockCaseRepo.On("GetAllCases").Return(cases, nil)

	// Act
	returnedCases, err := suite.caseUseCase.GetAllCases()

	// Assert
	suite.Nil(err)
	suite.Len(returnedCases, 1)
	suite.Equal("decrypted_field", returnedCases[0].Title)
}

func (suite *CaseUseCaseTestSuite) TestGetAllCases_Fail() {
	// Expect case repo to return an error
	suite.mockCaseRepo.On("GetAllCases").Return(nil, domain.ErrCaseFetchFailed)

	// Act
	returnedCases, err := suite.caseUseCase.GetAllCases()

	// Assert
	suite.Equal(domain.ErrCaseFetchFailed, err)
	suite.Empty(returnedCases)
}

func (suite *CaseUseCaseTestSuite) TestGetCaseByID_Success() {
	// Arrange
	caseID := uuid.New()
	caseModel := &domain.Case{
		Title: "Test Case",
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil)

	// Expect case repo to return case successfully
	suite.mockCaseRepo.On("GetCaseByID", caseID).Return(caseModel, nil)

	// Act
	returnedCase, err := suite.caseUseCase.GetCaseByID(caseID)

	// Assert
	suite.Nil(err)
	suite.Equal("decrypted_field", returnedCase.Title)
}

func (suite *CaseUseCaseTestSuite) TestGetCaseByID_Fail() {
	// Arrange
	caseID := uuid.New()

	// Expect case repo to return an error
	suite.mockCaseRepo.On("GetCaseByID", caseID).Return(nil, domain.ErrCaseNotFound)

	// Act
	returnedCase, err := suite.caseUseCase.GetCaseByID(caseID)

	// Assert
	suite.Equal(domain.ErrCaseNotFound, err)
	suite.Nil(returnedCase)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesBySubmitterID_Success() {
	// Arrange
	submitterID := uuid.New()
	cases := []*domain.Case{
		{
			Title: "Test Case",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil)

	// Expect case repo to return cases successfully
	suite.mockCaseRepo.On("GetCasesBySubmitterID", submitterID).Return(cases, nil)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesBySubmitterID(submitterID)

	// Assert
	suite.Nil(err)
	suite.Len(returnedCases, 1)
	suite.Equal("decrypted_field", returnedCases[0].Title)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesBySubmitterID_Fail() {
	// Arrange
	submitterID := uuid.New()

	// Expect case repo to return an error
	suite.mockCaseRepo.On("GetCasesBySubmitterID", submitterID).Return(nil, domain.ErrCaseFetchFailed)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesBySubmitterID(submitterID)

	// Assert
	suite.Equal(domain.ErrCaseFetchFailed, err)
	suite.Empty(returnedCases)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesByCounselorID_Success() {
	// Arrange
	counselorID := uuid.New()
	cases := []*domain.Case{
		{
			Title: "Test Case",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil)

	// Expect case repo to return cases successfully
	suite.mockCaseRepo.On("GetCasesByCounselorID", counselorID).Return(cases, nil)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesByCounselorID(counselorID)

	// Assert
	suite.Nil(err)
	suite.Len(returnedCases, 1)
	suite.Equal("decrypted_field", returnedCases[0].Title)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesByCounselorID_Fail() {
	// Arrange
	counselorID := uuid.New()

	// Expect case repo to return an error
	suite.mockCaseRepo.On("GetCasesByCounselorID", counselorID).Return(nil, domain.ErrCaseFetchFailed)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesByCounselorID(counselorID)

	// Assert
	suite.Equal(domain.ErrCaseFetchFailed, err)
	suite.Empty(returnedCases)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesByStatus_Success() {
	// Arrange
	status := "pending"
	cases := []*domain.Case{
		{
			Title: "Test Case",
		},
	}

	// Expect encryption service to decrypt all fields successfully
	suite.mockEncryptService.On("Decrypt", mock.Anything).Return("decrypted_field", nil)

	// Expect case repo to return cases successfully
	suite.mockCaseRepo.On("GetCasesByStatus", status).Return(cases, nil)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesByStatus(status)

	// Assert
	suite.Nil(err)
	suite.Len(returnedCases, 1)
	suite.Equal("decrypted_field", returnedCases[0].Title)
}

func (suite *CaseUseCaseTestSuite) TestGetCasesByStatus_Fail() {
	// Arrange
	status := "pending"

	// Expect case repo to return an error
	suite.mockCaseRepo.On("GetCasesByStatus", status).Return(nil, domain.ErrCaseFetchFailed)

	// Act
	returnedCases, err := suite.caseUseCase.GetCasesByStatus(status)

	// Assert
	suite.Equal(domain.ErrCaseFetchFailed, err)
	suite.Empty(returnedCases)
}


func (suite *CaseUseCaseTestSuite) TestUpdateCase_Success() {
	// Arrange
	caseID := uuid.New()
	caseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	// Expect encryption service to encrypt all fields successfully
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted_field", nil).Times(3)

	// Expect case repo to update case successfully
	suite.mockCaseRepo.On("UpdateCaseFields", caseID, mock.Anything).Return(nil)

	// Act
	err := suite.caseUseCase.UpdateCase(caseID, caseDto)

	// Assert
	suite.Nil(err)
}

func (suite *CaseUseCaseTestSuite) TestUpdateCase_Fail() {
	// Arrange
	caseID := uuid.New()
	caseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	// Expect encryption service to encrypt all fields successfully
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted_field", nil).Times(3)

	// Expect case repo to return an error
	suite.mockCaseRepo.On("UpdateCaseFields", caseID, mock.Anything).Return(domain.ErrCaseUpdateFailed)

	// Act
	err := suite.caseUseCase.UpdateCase(caseID, caseDto)

	// Assert
	suite.Equal(domain.ErrCaseUpdateFailed, err)
}
func TestCaseUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CaseUseCaseTestSuite))
}
