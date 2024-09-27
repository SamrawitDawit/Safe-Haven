package usecases

import (
	"backend/domain"
	"backend/usecases/dto"
	"backend/usecases/interfaces"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CaseUseCaseInterface interface {
	CreateCase(CaseDto dto.CaseDto) (*domain.Case, *domain.CustomError)
	UpdateCase(caseID uuid.UUID, CaseDto dto.CaseDto) *domain.CustomError
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

func validateCaseDto(CaseDto dto.CaseDto) *domain.CustomError {
	if CaseDto.Description == "" && CaseDto.ImageURL == "" {
		return domain.ErrIncompleteCaseInformation
	}
	return nil
}

func (r *CaseUseCase) decrypt(Cases []*domain.Case) ([]*domain.Case, *domain.CustomError) {
	for _, Case := range Cases {
		if Case.Title != "" {
			encryptedTitle, err := r.EncrypService.Decrypt(Case.Title)
			if err != nil {
				return nil, err
			}
			Case.Title = encryptedTitle
		}
		if Case.Description != "" {
			encryptedDesc, err := r.EncrypService.Decrypt(Case.Description)
			if err != nil {
				return nil, err
			}
			Case.Description = encryptedDesc
		}
		if Case.ImageURL != "" {
			encryptedURL, err := r.EncrypService.Decrypt(Case.ImageURL)
			if err != nil {
				return nil, err
			}
			Case.ImageURL = encryptedURL
		}
		if Case.VideoURL != "" {
			encryptedURL, err := r.EncrypService.Decrypt(Case.VideoURL)
			if err != nil {
				return nil, err
			}
			Case.VideoURL = encryptedURL
		}
		if Case.Location != "" {
			encryptedLoc, err := r.EncrypService.Decrypt(Case.Location)
			if err != nil {
				return nil, err
			}
			Case.Location = encryptedLoc
		}
	}
	return Cases, nil
}

func (r *CaseUseCase) CreateCase(CaseDto dto.CaseDto) (*domain.Case, *domain.CustomError) {
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
		VideoURL:          CaseDto.VideoURL,
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
	if CaseDto.VideoURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(CaseDto.VideoURL)
		if err != nil {
			return nil, err
		}
		new_Case.VideoURL = encryptedURL
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
	decryptedCases, err := r.decrypt(Cases)
	if err != nil {
		return nil, err
	}
	return decryptedCases, nil
}

func (r *CaseUseCase) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCaseByID(CaseID)
	if err != nil {
		return nil, err
	}
	decryptedCases, err := r.decrypt([]*domain.Case{Case})
	if err != nil {
		return nil, err
	}
	return decryptedCases[0], nil
}

func (r *CaseUseCase) GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesByCounselorID(counselorID)
	if err != nil {
		return nil, err
	}
	decryptedCases, err := r.decrypt(Case)
	if err != nil {
		return nil, err
	}
	return decryptedCases, nil
}

func (r *CaseUseCase) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesBySubmitterID(SubmitterID)
	if err != nil {
		return nil, err
	}
	decryptedCases, err := r.decrypt(Case)
	if err != nil {
		return nil, err
	}
	return decryptedCases, nil
}

func (r *CaseUseCase) GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError) {
	Case, err := r.CaseRepo.GetCasesByStatus(status)
	if err != nil {
		return nil, err
	}
	decryptedCases, err := r.decrypt(Case)
	if err != nil {
		return nil, err
	}
	return decryptedCases, nil
}

func (r *CaseUseCase) UpdateCase(caseID uuid.UUID, CaseDto dto.CaseDto) *domain.CustomError {
	fmt.Println("case Id", caseID)
	updatedFields := map[string]interface{}{}
	if CaseDto.Title != "" {
		encryptedTitle, err := r.EncrypService.Encrypt(CaseDto.Title)
		if err != nil {
			return err
		}
		updatedFields["title"] = encryptedTitle
	}
	if CaseDto.Description != "" {
		encryptedDesc, err := r.EncrypService.Encrypt(CaseDto.Description)
		if err != nil {
			return err
		}
		updatedFields["description"] = encryptedDesc
	}
	if CaseDto.ImageURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(CaseDto.ImageURL)
		if err != nil {
			return err
		}
		updatedFields["image_url"] = encryptedURL
	}
	if CaseDto.VideoURL != "" {
		encryptedURL, err := r.EncrypService.Encrypt(CaseDto.VideoURL)
		if err != nil {
			return err
		}
		updatedFields["video_url"] = encryptedURL
	}
	if CaseDto.Location != "" {
		encryptedLoc, err := r.EncrypService.Encrypt(CaseDto.Location)
		if err != nil {
			return err
		}
		updatedFields["location"] = encryptedLoc
	}
	err := r.CaseRepo.UpdateCaseFields(caseID, updatedFields)
	if err != nil {
		return err
	}
	return nil
}