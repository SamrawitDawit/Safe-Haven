package interfaces

import (
	"backend/domain"

	"github.com/google/uuid"
)

type CaseRepositoryInterface interface {
	CreateCase(Case *domain.Case) *domain.CustomError
	UpdateCaseFields(CaseID uuid.UUID, fields map[string]interface{}) *domain.CustomError
	GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError)
	GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError)
	GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError)
	GetAllCases() ([]*domain.Case, *domain.CustomError)
	GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError)
	DeleteCase(CaseID uuid.UUID) *domain.CustomError
}
