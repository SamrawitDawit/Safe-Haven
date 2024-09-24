package controllers

import (
	"backend/usecases"
	"backend/usecases/dto"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReportController struct {
	ReportUsecase usecases.ReportUseCaseInterface
}

func NewReportController(reportUsecase usecases.ReportUseCaseInterface) *ReportController {
	return &ReportController{
		ReportUsecase: reportUsecase,
	}
}

func (ctrl *ReportController) CreateReport(c *gin.Context) {
	var createReportDto dto.CreateReportDto
	if err := c.BindJSON(&createReportDto); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	created_report, cerr := ctrl.ReportUsecase.CreateReport(createReportDto)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report submission failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "Report submitted successfully", created_report))
}

func (ctrl *ReportController) UpdateReport(c *gin.Context) {
	var updateReportDto dto.UpdateReportDto
	if err := c.BindJSON(&updateReportDto); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	cerr := ctrl.ReportUsecase.UpdateReport(updateReportDto)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report Update Failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report Update successful", nil))
}
func (ctrl *ReportController) GetReportByID(c *gin.Context) {
	ReportID, err := uuid.Parse(c.Param("report_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	report, cerr := ctrl.ReportUsecase.GetReportByID(ReportID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report fetch successful", report))
}

func (ctrl *ReportController) GetReportsByReporterID(c *gin.Context) {
	ReporterID, err := uuid.Parse(c.Param("reporter_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	report, cerr := ctrl.ReportUsecase.GetReportsByReporterID(ReporterID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report fetch successful", report))
}

func (ctrl *ReportController) GetReportsByCounselorID(c *gin.Context) {
	CounselorID, err := uuid.Parse(c.Param("counselor_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	report, cerr := ctrl.ReportUsecase.GetReportsByCounselorID(CounselorID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report fetch successful", report))
}

func (ctrl *ReportController) GetReportsByStatus(c *gin.Context) {
	status := c.Param("status")
	report, cerr := ctrl.ReportUsecase.GetReportsByStatus(status)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report fetch successful", report))
}

func (ctrl *ReportController) GetAllReports(c *gin.Context) {
	report, cerr := ctrl.ReportUsecase.GetAllReports()
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report fetch failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report fetch successful", report))
}

func (ctrl *ReportController) DeleteReport(c *gin.Context) {
	ReportID, err := uuid.Parse(c.Param("report_id"))
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	cerr := ctrl.ReportUsecase.DeleteReport(ReportID)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Report deletion failed", cerr.Message)
		c.JSON(cerr.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Report deletion successful", nil))
}
