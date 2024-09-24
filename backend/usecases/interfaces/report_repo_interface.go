package interfaces

import (
	"backend/domain"

	"github.com/google/uuid"
)

type ReportRepositoryInterface interface {
	CreateReport(report *domain.Report) *domain.CustomError
	UpdateReportFields(reportID uuid.UUID, fields map[string]interface{}) *domain.CustomError
	GetReportByID(reportID uuid.UUID) (*domain.Report, *domain.CustomError)
	GetReportsByReporterID(reporterID uuid.UUID) ([]*domain.Report, *domain.CustomError)
	GetReportsByCounselorID(counselorID uuid.UUID) ([]*domain.Report, *domain.CustomError)
	GetAllReports() ([]*domain.Report, *domain.CustomError)
	GetReportsByStatus(status string) ([]*domain.Report, *domain.CustomError)
	DeleteReport(reportID uuid.UUID) *domain.CustomError
}
