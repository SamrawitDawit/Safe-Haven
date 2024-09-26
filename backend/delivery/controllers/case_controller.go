package controllers

import (
	"backend/usecases"
	"backend/usecases/dto"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CaseController struct {
	CaseUsecase usecases.CaseUseCaseInterface
}

func NewCaseController(CaseUsecase usecases.CaseUseCaseInterface) *CaseController {
	return &CaseController{
		CaseUsecase: CaseUsecase,
	}
}

func (ctrl *CaseController) CreateCase(c *gin.Context) {
	var createCaseDto dto.CreateCaseDto
	if err := c.BindJSON(&createCaseDto); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	created_Case, cerr := ctrl.CaseUsecase.CreateCase(createCaseDto)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case submission failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "Case submitted successfully", created_Case))
}

func (ctrl *CaseController) UpdateCase(c *gin.Context) {
	var updateCaseDto dto.UpdateCaseDto
	if err := c.BindJSON(&updateCaseDto); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	cerr := ctrl.CaseUsecase.UpdateCase(updateCaseDto)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case Update Failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case Update successful", nil))
}
func (ctrl *CaseController) GetCaseByID(c *gin.Context) {
	CaseID, err := uuid.Parse(c.Param("Case_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Case, cerr := ctrl.CaseUsecase.GetCaseByID(CaseID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case fetch successful", Case))
}

func (ctrl *CaseController) GetCasesBySubmitterID(c *gin.Context) {
	SubmitterID, err := uuid.Parse(c.Param("submitter_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Case, cerr := ctrl.CaseUsecase.GetCasesBySubmitterID(SubmitterID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case fetch successful", Case))
}

func (ctrl *CaseController) GetCasesByCounselorID(c *gin.Context) {
	CounselorID, err := uuid.Parse(c.Param("counselor_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Case, cerr := ctrl.CaseUsecase.GetCasesByCounselorID(CounselorID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case fetch successful", Case))
}

func (ctrl *CaseController) GetCasesByStatus(c *gin.Context) {
	status := c.Param("status")
	Case, cerr := ctrl.CaseUsecase.GetCasesByStatus(status)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case fetch successful", Case))
}

func (ctrl *CaseController) GetAllCases(c *gin.Context) {
	Case, cerr := ctrl.CaseUsecase.GetAllCases()
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case fetch successful", Case))
}

func (ctrl *CaseController) DeleteCase(c *gin.Context) {
	CaseID, err := uuid.Parse(c.Param("Case_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	cerr := ctrl.CaseUsecase.DeleteCase(CaseID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Case deletion failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Case deletion successful", nil))
}
