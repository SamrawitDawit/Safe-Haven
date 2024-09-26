package usecases

import (
	"backend/domain"
	"backend/usecases/dto"
	"backend/usecases/interfaces"
	"time"

	"github.com/google/uuid"
)

type CaseUseCaseInterface interface {
	CreateCase(CaseDto dto.CreateCaseDto) (*domain.Case, *domain.CustomError)
	UpdateCase(CaseDto dto.UpdateCaseDto) *domain.CustomError
	GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError)
	GetAllCases() ([]*domain.Case, *domain.CustomError)
	GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError)
	GetCasesByCounselorID(counselor uuid.UUID) ([]*domain.Case, *domain.CustomError)
	GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError)
	DeleteCase(CaseID uuid.UUID) *domain.CustomError
}
type CaseUseCase struct {
	CaseRepo      interfaces.CaseRepositoryInterface
	EncrypService interfaces.EncryptionServiceInterface
}

func NewCaseUseCase(CaseRepo interfaces.CaseRepositoryInterface, encryptService interfaces.EncryptionServiceInterface) CaseUseCaseInterface {
	return &CaseUseCase{
		CaseRepo:      CaseRepo,
		EncrypService: encryptService,
	}
}

func validateCaseDto(CaseDto dto.CreateCaseDto) *domain.CustomError {
	if CaseDto.Description == "" || CaseDto.ImageURL == "" {
		return domain.ErrIncompleteCaseInformation
	}
	return nil
}

func (r *CaseUseCase) CreateCase(CaseDto dto.CreateCaseDto) (*domain.Case, *domain.CustomError) {
	err := validateCaseDto(CaseDto)
	if err != nil {
		return nil, err
	}
	new_Case := &domain.Case{
		ID:                uuid.New(),
		SubmitterID:       CaseDto.SubmitterID,
		Title:             CaseDto.Title,
		Description:       CaseDto.Description,
		ImageURL:          CaseDto.ImageURL,
		Location:          CaseDto.Location,
		Status:            "pending",
		SubmittedAt:       time.Now(),
		CounselorAssigned: false,
	}
	if CaseDto.Title != "" {
		encryptedTitle, err := r.EncrypService.Encrypt(CaseDto.Title)
		if err != nil {
			return nil, err
		}
		new_Case.Title = encryptedTitle
	}
	if CaseDto.Description != "" {
		encryptedDesc, err := r.EncrypService.Encrypt(CaseDto.Description)
		if err != nil {
			return nil, err
		}
		new_Case.Description = encryptedDesc
	}
	if CaseDto.ImageURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(CaseDto.ImageURL)
		if err != nil {
			return nil, err
		}
		new_Case.ImageURL = encryptedURL
	}
	if CaseDto.Location != "" {
		encryptedLoc, err := r.EncrypService.Encrypt(CaseDto.Location)
		if err != nil {
			return nil, err
		}
		new_Case.Location = encryptedLoc
	}
	err = r.CaseRepo.CreateCase(new_Case)
	if err != nil {
		return nil, err
	}
	return new_Case, err
}

func (r *CaseUseCase) DeleteCase(CaseID uuid.UUID) *domain.CustomError {
	err := r.CaseRepo.DeleteCase(CaseID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CaseUseCase) GetAllCases() ([]*domain.Case, *domain.CustomError) {
	Cases, err := r.CaseRepo.GetAllCases()
	if err != nil {
		return nil, err
	}
	return Cases, nil
}

func (r *CaseUseCase) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCaseByID(CaseID)
	if err != nil {
		return nil, err
	}
	return Case, nil
}

func (r *CaseUseCase) GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesByCounselorID(counselorID)
	if err != nil {
		return nil, err
	}
	return Case, nil
}

func (r *CaseUseCase) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesBySubmitterID(SubmitterID)
	if err != nil {
		return nil, err
	}
	return Case, nil
}

func (r *CaseUseCase) GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesByStatus(status)
	if err != nil {
		return nil, err
	}
	return Case, nil
}

func (r *CaseUseCase) UpdateCase(CaseDto dto.UpdateCaseDto) *domain.CustomError {
	updatedFields := map[string]interface{}{}
	if CaseDto.Title != "" {
		encryptedTitle, err := r.EncrypService.Encrypt(CaseDto.Title)
		if err != nil {
			return err
		}
		updatedFields["Title"] = encryptedTitle
	}
	if CaseDto.Description != "" {
		encryptedDesc, err := r.EncrypService.Encrypt(CaseDto.Description)
		if err != nil {
			return err
		}
		updatedFields["Description"] = encryptedDesc
	}
	if CaseDto.ImageURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(CaseDto.ImageURL)
		if err != nil {
			return err
		}
		updatedFields["ImageURL"] = encryptedURL
	}
	if CaseDto.Location != "" {
		encryptedLoc, err := r.EncrypService.Encrypt(CaseDto.Location)
		if err != nil {
			return err
		}
		updatedFields["Location"] = encryptedLoc
	}
	err := r.CaseRepo.UpdateCaseFields(CaseDto.ID, updatedFields)
	if err != nil {
		return err
	}
	return nil
}
