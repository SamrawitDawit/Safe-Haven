package usecases

import (
	"backend/domain"
	"backend/usecases/dto"
	"backend/usecases/interfaces"
	"time"

	"github.com/google/uuid"
)

type ReportUseCaseInterface interface {
	CreateReport(reportDto dto.CreateReportDto) (*domain.Report, *domain.CustomError)
	UpdateReport(reportDto dto.UpdateReportDto) *domain.CustomError
	GetReportByID(reportID uuid.UUID) (*domain.Report, *domain.CustomError)
	GetAllReports() ([]*domain.Report, *domain.CustomError)
	GetReportsByReporterID(reporterID uuid.UUID) ([]*domain.Report, *domain.CustomError)
	GetReportsByCounselorID(counselor uuid.UUID) ([]*domain.Report, *domain.CustomError)
	GetReportsByStatus(status string) ([]*domain.Report, *domain.CustomError)
	DeleteReport(reportID uuid.UUID) *domain.CustomError
}
type ReportUseCase struct {
	ReportRepo    interfaces.ReportRepositoryInterface
	EncrypService interfaces.EncryptionServiceInterface
}

func NewReportUseCase(reportRepo interfaces.ReportRepositoryInterface, encryptService interfaces.EncryptionServiceInterface) ReportUseCaseInterface {
	return &ReportUseCase{
		ReportRepo:    reportRepo,
		EncrypService: encryptService,
	}
}

func validateReportDto(reportDto dto.CreateReportDto) *domain.CustomError {
	if reportDto.Description == "" || reportDto.ImageURL == "" {
		return domain.ErrIncompleteReportInformation
	}
	return nil
}

func (r *ReportUseCase) CreateReport(reportDto dto.CreateReportDto) (*domain.Report, *domain.CustomError) {
	err := validateReportDto(reportDto)
	if err != nil {
		return nil, err
	}
	new_report := &domain.Report{
		ID:                uuid.New(),
		ReporterID:        reportDto.ReporterID,
		Title:             reportDto.Title,
		Description:       reportDto.Description,
		ImageURL:          reportDto.ImageURL,
		Location:          reportDto.Location,
		Status:            "pending",
		ReportedAt:        time.Now(),
		CounselorAssigned: false,
	}
	if reportDto.Title != "" {
		encryptedTitle, err := r.EncrypService.Encrypt(reportDto.Title)
		if err != nil {
			return nil, err
		}
		new_report.Title = encryptedTitle
	}
	if reportDto.Description != "" {
		encryptedDesc, err := r.EncrypService.Encrypt(reportDto.Description)
		if err != nil {
			return nil, err
		}
		new_report.Description = encryptedDesc
	}
	if reportDto.ImageURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(reportDto.ImageURL)
		if err != nil {
			return nil, err
		}
		new_report.ImageURL = encryptedURL
	}
	if reportDto.Location != "" {
		encryptedLoc, err := r.EncrypService.Encrypt(reportDto.Location)
		if err != nil {
			return nil, err
		}
		new_report.Location = encryptedLoc
	}
	err = r.ReportRepo.CreateReport(new_report)
	if err != nil {
		return nil, err
	}
	return new_report, err
}

func (r *ReportUseCase) DeleteReport(reportID uuid.UUID) *domain.CustomError {
	err := r.ReportRepo.DeleteReport(reportID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReportUseCase) GetAllReports() ([]*domain.Report, *domain.CustomError) {
	reports, err := r.ReportRepo.GetAllReports()
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *ReportUseCase) GetReportByID(reportID uuid.UUID) (*domain.Report, *domain.CustomError) {
	report, err := r.ReportRepo.GetReportByID(reportID)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *ReportUseCase) GetReportsByCounselorID(counselorID uuid.UUID) ([]*domain.Report, *domain.CustomError) {
	report, err := r.ReportRepo.GetReportsByCounselorID(counselorID)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *ReportUseCase) GetReportsByReporterID(reporterID uuid.UUID) ([]*domain.Report, *domain.CustomError) {
	report, err := r.ReportRepo.GetReportsByReporterID(reporterID)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *ReportUseCase) GetReportsByStatus(status string) ([]*domain.Report, *domain.CustomError) {
	report, err := r.ReportRepo.GetReportsByStatus(status)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *ReportUseCase) UpdateReport(reportDto dto.UpdateReportDto) *domain.CustomError {
	updatedFields := map[string]interface{}{}
	if reportDto.Title != "" {
		encryptedTitle, err := r.EncrypService.Encrypt(reportDto.Title)
		if err != nil {
			return err
		}
		updatedFields["Title"] = encryptedTitle
	}
	if reportDto.Description != "" {
		encryptedDesc, err := r.EncrypService.Encrypt(reportDto.Description)
		if err != nil {
			return err
		}
		updatedFields["Description"] = encryptedDesc
	}
	if reportDto.ImageURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(reportDto.ImageURL)
		if err != nil {
			return err
		}
		updatedFields["ImageURL"] = encryptedURL
	}
	if reportDto.Location != "" {
		encryptedLoc, err := r.EncrypService.Encrypt(reportDto.Location)
		if err != nil {
			return err
		}
		updatedFields["Location"] = encryptedLoc
	}
	err := r.ReportRepo.UpdateReportFields(reportDto.ID, updatedFields)
	if err != nil {
		return err
	}
	return nil
}
