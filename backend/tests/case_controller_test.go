package tests

import (
	"backend/delivery/controllers"
	"backend/domain"
	"backend/tests/mocks"
	"backend/usecases/dto"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CaseControllerTestSuite struct {
	suite.Suite
	mockUsecase   *mocks.CaseUseCaseInterface
	mockRecaptcha *mocks.RecaptchaInterface
	controller    *controllers.CaseController
	recorder      *httptest.ResponseRecorder
}

func (suite *CaseControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.CaseUseCaseInterface)
	suite.mockRecaptcha = new(mocks.RecaptchaInterface)
	suite.recorder = httptest.NewRecorder()
	suite.controller = controllers.NewCaseController(suite.mockUsecase, suite.mockRecaptcha)
}

func (suite *CaseControllerTestSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *CaseControllerTestSuite) TestCreateCase_Success() {
	// Arrange

	createCaseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	createdCase := &domain.Case{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	suite.mockRecaptcha.On("CreateAssessment", "test_token").Return(float32(0.5), nil)
	suite.mockUsecase.On("CreateCase", createCaseDto).Return(createdCase, nil)
	body, _ := json.Marshal(createCaseDto)
	req, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("recaptcha-token", "test_token")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	suite.controller.CreateCase(c)

	suite.Equal(http.StatusCreated, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case submitted successfully")
}

func (suite *CaseControllerTestSuite) TestCreateCase_Fail() {
	// Arrange
	createCaseDto := dto.CaseDto{
		SubmitterID: uuid.New(),
	}
	suite.mockRecaptcha.On("CreateAssessment", "test_token").Return(float32(0.5), nil)
	suite.mockUsecase.On("CreateCase", createCaseDto).Return(nil, domain.ErrIncompleteCaseInformation)
	body, _ := json.Marshal(createCaseDto)
	req, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("recaptcha-token", "test_token")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.CreateCase(c)

	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case submission failed")
}

func (suite *CaseControllerTestSuite) TestUpdateCase_Success() {
	// Arrange
	caseID := uuid.New()
	updateCaseDto := dto.CaseDto{
		Title: "Test Case",
	}

	suite.mockUsecase.On("UpdateCase", caseID, updateCaseDto).Return(nil)
	body, _ := json.Marshal(updateCaseDto)
	req, _ := http.NewRequest(http.MethodPut, "/update/case_id:"+caseID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.UpdateCase(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case Update successful")
}

func (suite *CaseControllerTestSuite) TestUpdateCase_Fail() {
	// Arrange
	caseID := uuid.New()
	updateCaseDto := dto.CaseDto{
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	suite.mockUsecase.On("UpdateCase", caseID, updateCaseDto).Return(domain.ErrCaseUpdateFailed)
	body, _ := json.Marshal(updateCaseDto)
	req, _ := http.NewRequest(http.MethodPut, "/update/"+caseID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.UpdateCase(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case Update Failed")
}

func (suite *CaseControllerTestSuite) TestGetCaseByID_Success() {
	// Arrange
	caseID := uuid.New()
	caseData := &domain.Case{
		ID:          caseID,
		Title:       "Test Case",
		Description: "This is a test case",
		ImageURL:    "https://example.com/image.jpg",
		SubmitterID: uuid.New(),
	}

	suite.mockUsecase.On("GetCaseByID", caseID).Return(caseData, nil)
	req, _ := http.NewRequest(http.MethodGet, "/case/"+caseID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.GetCaseByID(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch successful")
}

func (suite *CaseControllerTestSuite) TestGetCaseByID_Fail() {
	// Arrange
	caseID := uuid.New()

	suite.mockUsecase.On("GetCaseByID", caseID).Return(nil, domain.ErrCaseFetchFailed)
	req, _ := http.NewRequest(http.MethodGet, "/case/"+caseID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.GetCaseByID(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch failed")
}

func (suite *CaseControllerTestSuite) TestGetCasesBySubmitterID_Success() {
	// Arrange
	submitterID := uuid.New()
	caseData := []*domain.Case{
		{
			ID:          uuid.New(),
			Title:       "Test Case",
			Description: "This is a test case",
			ImageURL:    "https://example.com/image.jpg",
			SubmitterID: submitterID,
		},
	}

	suite.mockUsecase.On("GetCasesBySubmitterID", submitterID).Return(caseData, nil)
	req, _ := http.NewRequest(http.MethodGet, "/submitter/"+submitterID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "submitter_id", Value: submitterID.String()},
	}

	suite.controller.GetCasesBySubmitterID(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch successful")
}

func (suite *CaseControllerTestSuite) TestGetCasesBySubmitterID_Fail() {
	// Arrange
	submitterID := uuid.New()

	suite.mockUsecase.On("GetCasesBySubmitterID", submitterID).Return(nil, domain.ErrCaseFetchFailed)
	req, _ := http.NewRequest(http.MethodGet, "/submitter/"+submitterID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "submitter_id", Value: submitterID.String()},
	}

	suite.controller.GetCasesBySubmitterID(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch failed")
}

func (suite *CaseControllerTestSuite) TestGetCasesByCounselorID_Success() {

	// Arrange
	counselorID := uuid.New()
	caseData := []*domain.Case{
		{
			ID:          uuid.New(),
			Title:       "Test Case",
			Description: "This is a test case",
			ImageURL:    "https://example.com/image.jpg",
			SubmitterID: uuid.New(),
		},
	}

	suite.mockUsecase.On("GetCasesByCounselorID", counselorID).Return(caseData, nil)
	req, _ := http.NewRequest(http.MethodGet, "/counselor/"+counselorID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "counselor_id", Value: counselorID.String()},
	}

	suite.controller.GetCasesByCounselorID(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch successful")
}

func (suite *CaseControllerTestSuite) TestGetCasesByCounselorID_Fail() {
	// Arrange
	counselorID := uuid.New()

	suite.mockUsecase.On("GetCasesByCounselorID", counselorID).Return(nil, domain.ErrCaseFetchFailed)
	req, _ := http.NewRequest(http.MethodGet, "/counselor/"+counselorID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "counselor_id", Value: counselorID.String()},
	}

	suite.controller.GetCasesByCounselorID(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch failed")
}

func (suite *CaseControllerTestSuite) TestGetCasesByStatus_Success() {
	// Arrange
	status := "pending"
	caseData := []*domain.Case{
		{
			ID:          uuid.New(),
			Title:       "Test Case",
			Description: "This is a test case",
			ImageURL:    "https://example.com/image.jpg",
			SubmitterID: uuid.New(),
		},
	}

	suite.mockUsecase.On("GetCasesByStatus", status).Return(caseData, nil)
	req, _ := http.NewRequest(http.MethodGet, "/status/"+status, nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "status", Value: status},
	}

	suite.controller.GetCasesByStatus(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch successful")
}

func (suite *CaseControllerTestSuite) TestGetCasesByStatus_Fail() {
	// Arrange
	status := "pending"

	suite.mockUsecase.On("GetCasesByStatus", status).Return(nil, domain.ErrCaseFetchFailed)
	req, _ := http.NewRequest(http.MethodGet, "/status/"+status, nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "status", Value: status},
	}

	suite.controller.GetCasesByStatus(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch failed")
}

func (suite *CaseControllerTestSuite) TestGetAllCases_Success() {
	// Arrange
	caseData := []*domain.Case{
		{
			ID:          uuid.New(),
			Title:       "Test Case",
			Description: "This is a test case",
			ImageURL:    "https://example.com/image.jpg",
			SubmitterID: uuid.New(),
		},
	}

	suite.mockUsecase.On("GetAllCases").Return(caseData, nil)
	req, _ := http.NewRequest(http.MethodGet, "/all", nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.GetAllCases(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch successful")
}

func (suite *CaseControllerTestSuite) TestGetAllCases_Fail() {
	// Arrange
	suite.mockUsecase.On("GetAllCases").Return(nil, domain.ErrCaseFetchFailed)
	req, _ := http.NewRequest(http.MethodGet, "/all", nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.GetAllCases(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case fetch failed")
}

func (suite *CaseControllerTestSuite) TestDeleteCase_Success() {
	// Arrange
	caseID := uuid.New()

	suite.mockUsecase.On("DeleteCase", caseID).Return(nil)
	req, _ := http.NewRequest(http.MethodDelete, "/delete/"+caseID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.DeleteCase(c)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case deletion successful")
}

func (suite *CaseControllerTestSuite) TestDeleteCase_Fail() {
	// Arrange
	caseID := uuid.New()

	suite.mockUsecase.On("DeleteCase", caseID).Return(domain.ErrCaseDeletionFailed)
	req, _ := http.NewRequest(http.MethodDelete, "/delete/"+caseID.String(), nil)

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{Key: "case_id", Value: caseID.String()},
	}

	suite.controller.DeleteCase(c)

	suite.Equal(http.StatusInternalServerError, suite.recorder.Code)
	suite.Contains(suite.recorder.Body.String(), "Case deletion failed")
}

func TestCaseControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CaseControllerTestSuite))
}
