package usecases

import (
	"backend/domain"
	"backend/infrastructure"
	"backend/repositories"
	"backend/usecases/dto"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type AuthUseCaseInterface interface {
	Register(registerDTO dto.RegisterDTO) error
	Login(loginDTO dto.LoginDTO) (*domain.User, error)
	AnonymousLogin(anonUserDTO dto.LoginDTO) (*domain.User, error)
	RefreshToken(refreshToken string) (string, string, error)
}

func NewAuthUseCase(userRepo repositories.UserRepository, anonUserRepo repositories.UserRepository, jwtservice infrastructure.JWTService, emailService infrastructure.EmailService, pwdService infrastructure.HashingService) AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo:     userRepo,
		jwtService:   jwtservice,
		emailService: emailService,
		pwdService:   pwdService,
	}
}

type AuthUseCase struct {
	userRepo     repositories.UserRepository
	jwtService   infrastructure.JWTService
	emailService infrastructure.EmailService
	pwdService   infrastructure.HashingService
}

func (a *AuthUseCase) AnonymousLogin(anonUserDTO dto.LoginDTO) (*domain.User, error) {
	return a.userRepo.GetUserByAnonymousDifferentiator(anonUserDTO.AnonymousDifferenitator)
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
		UserType:          registerDTO.UserType,
		Role:              "regular",
		Active:            true,
		Verified:          false,
		CounselorAssigned: false,
	}
	return a.userRepo.CreateUser(user)
}
func (a *AuthUseCase) RefreshToken(refreshToken string) (string, string, error) {
	token, err := a.jwtService.ValidateToken(refreshToken)
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}
	claims, err := a.jwtService.ExtractTokenClaims(token)
	if err != nil {
		return "", "", errors.New("invalid token claims")
	}

	id := claims["id"].(string)
	user, err := a.userRepo.GetUserByID(id)
	if err != nil || user == nil {
		return "", "", err
	}

	token, err = a.jwtService.ValidateToken(user.RefreshToken)
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}
	if user.RefreshToken != refreshToken {
		return "", "", errors.New("invalid token")
	}

	accessToken, refreshToken, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	err = a.userRepo.UpdateUser(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthUseCase) ForgotPassword(email string) error {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	// send email
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min, max := 10000, 100000

	randomNumber := r.Int63n(int64(max-min)) + int64(min)
	resetToken, err := a.jwtService.GenerateResetToken(user.Email, randomNumber)
	if err != nil {
		return err
	}

	user.ResetCode = randomNumber
	user.ResetToken = resetToken
	err = a.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return a.emailService.SendResetPasswordEmail(user.Email, resetToken)
}

func (a *AuthUseCase) ResetPassword(token string, newPassword string) error {
	validatedToken, err := a.jwtService.ValidateToken(token)
	if err != nil || !validatedToken.Valid {
		return errors.New("invalid token")
	}
	claims, err := a.jwtService.ExtractTokenClaims(validatedToken)
	if err != nil {
		return errors.New("invalid token claims")
	}
	code, ok := claims["code"].(float64)
	if !ok {
		return errors.New("invalid token claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token claims")
	}
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return err
	}
	if user.ResetCode != int64(code) {
		return errors.New("invalid token")
	}
	if user.ResetToken != token {
		return errors.New("invalid token")
	}
	hashedPassword, err := a.pwdService.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.ResetCode = 0
	user.ResetToken = ""
	user.Password = string(hashedPassword)
	return a.userRepo.UpdateUser(user)
}

func (a *AuthUseCase) HandleGoogleCallback(user *domain.User) (string, string, error) {
	existingUser, err := a.userRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return "", "", err
	}
	if existingUser != nil {
		if !user.GoogleSignin {
			return "", "", errors.New("user already exists")
		}
		accessToken, refreshToken, err := a.jwtService.GenerateToken(user)
		if err != nil {
			return "", "", err
		}
		user.AccessToken = accessToken
		user.RefreshToken = refreshToken
		err = a.userRepo.UpdateUser(user)
		if err != nil {
			return "", "", err
		}
		return accessToken, refreshToken, nil

	}
	accessToken, refreshToken, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return "", "", err
	}
	user.ID = uuid.New()
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.Role = "regular"
	user.Active = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Verified = true
	user.CounselorAssigned = false
	user.GoogleSignin = true
	err = a.userRepo.CreateUser(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
