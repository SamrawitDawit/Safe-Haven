package usecases

import (
	"backend/domain"
	"backend/infrastructure"
	"backend/repositories"
	"backend/usecases/dto"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthUseCaseInterface interface {
	Register(registerDTO dto.RegisterDTO) error
	Login(loginDTO dto.LoginDTO) (*domain.User, error)
	AnonymousRegister(anonUserDTO dto.AnonymousUser) error
	AnonymousLogin(anonUserDTO dto.AnonymousUser) (*domain.AnonymousUser, error)
	RefreshToken(refreshToken string) (*dto.TokenResponseDto, error)
}

func NewAuthUseCase(userRepo repositories.UserRepository, anonUserRepo repositories.AnonymousUserRepository, jwtservice infrastructure.JWTService) AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo:     userRepo,
		anonUserRepo: anonUserRepo,
		jwtService:   &jwtservice,
	}
}

type AuthUseCase struct {
	userRepo     repositories.UserRepository
	anonUserRepo repositories.AnonymousUserRepository
	jwtService   *infrastructure.JWTService
}

func (a *AuthUseCase) AnonymousLogin(anonUserDTO dto.AnonymousUser) (*domain.AnonymousUser, error) {
	return a.anonUserRepo.GetUserByAnonymousDifferentiator(anonUserDTO.AnonymousDifferenitator)
}

func (a *AuthUseCase) AnonymousRegister(anonUserDTO dto.AnonymousUser) error {
	user := &domain.AnonymousUser{
		ID:                      uuid.New(),
		AnonymousDifferentiator: anonUserDTO.AnonymousDifferenitator,
		Password:                anonUserDTO.Password,
		CounselorAssigned:       false,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}
	return a.anonUserRepo.CreateAnonymousUser(user)
}

func (a *AuthUseCase) Login(loginDTO dto.LoginDTO) (*domain.User, error) {
	if loginDTO.Email != "" {
		user, err := a.userRepo.GetUserByEmail(loginDTO.Email)
		return user, err
	} else if loginDTO.PhoneNumber != "" {
		user, err := a.userRepo.GetUserByPhoneNumber(loginDTO.PhoneNumber)
		return user, err
	}
	return nil, errors.New("invalid login credentials")
}

func (a *AuthUseCase) Register(registerDTO dto.RegisterDTO) error {
	user := &domain.User{
		ID:                uuid.New(),
		FullName:          registerDTO.FullName,
		Email:             registerDTO.Email,
		Password:          registerDTO.Password,
		PhoneNumber:       registerDTO.PhoneNumber,
		Role:              "normal",
		Active:            true,
		Verified:          false,
		CounselorAssigned: false,
	}
	return a.userRepo.CreateNormalUser(user)
}
func (a *AuthUseCase) RefreshToken(refreshToken string) (*dto.TokenResponseDto, error) {
	panic("unimplemented")
}
