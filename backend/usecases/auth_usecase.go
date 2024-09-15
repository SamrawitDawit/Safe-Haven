package usecases

import (
	"backend/domain"
	"backend/usecases/dto"
	"backend/usecases/interfaces"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthUseCaseInterface interface {
	Register(registerDTO dto.RegisterDTO) error
	Login(loginDTO dto.LoginDTO) (string, string, error)
	RefreshToken(refreshToken string) (string, string, error)
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
	HandleGoogleCallback(user *domain.User) (string, string, error)
}

func NewAuthUseCase(userRepo interfaces.UserRepositoryInterface, jwtservice interfaces.JwtServiceInterface, emailService interfaces.EmailServiceInterface, pwdService interfaces.HashingServiceInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo:     userRepo,
		jwtService:   jwtservice,
		emailService: emailService,
		pwdService:   pwdService,
	}
}

type AuthUseCase struct {
	userRepo     interfaces.UserRepositoryInterface
	jwtService   interfaces.JwtServiceInterface
	emailService interfaces.EmailServiceInterface
	pwdService   interfaces.HashingServiceInterface
}

func (a *AuthUseCase) Login(loginDTO dto.LoginDTO) (string, string, error) {
	var user *domain.User
	if loginDTO.Email != "" {
		user, _ = a.userRepo.GetUserByEmail(loginDTO.Email)
		if user != nil && user.GoogleSignin {
			return "", "", errors.New("invalid login credentials")
		}
		if user == nil {
			return "", "", errors.New("invalid login credentials")
		}
	} else if loginDTO.PhoneNumber != "" {
		user, _ = a.userRepo.GetUserByPhoneNumber(loginDTO.PhoneNumber)
		if user == nil {
			return "", "", errors.New("invalid login credentials")
		}
	} else if loginDTO.AnonymousDifferentiator != "" {
		user, _ = a.userRepo.GetUserByAnonymousDifferentiator(loginDTO.AnonymousDifferentiator)
		if user == nil {
			return "", "", errors.New("invalid login credentials")
		}
	}
	err := a.pwdService.CheckPasswordHash(loginDTO.Password, user.Password)
	if err != nil {
		return "", "", errors.New("invalid login credentials")
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

func (a *AuthUseCase) Register(registerDTO dto.RegisterDTO) error {
	var existingUser *domain.User
	var err error
	//checking if user already exists
	if registerDTO.Email != "" {
		existingUser, err = a.userRepo.GetUserByEmail(registerDTO.Email)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if existingUser != nil && !existingUser.GoogleSignin {
			return errors.New("user already exists")
		}

	} else if registerDTO.PhoneNumber != "" {
		existingUser, err = a.userRepo.GetUserByPhoneNumber(registerDTO.PhoneNumber)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if existingUser != nil {
			return errors.New("user already exists")
		}
	} else if registerDTO.AnonymousDifferentiator != "" {
		existingUser, err = a.userRepo.GetUserByAnonymousDifferentiator(registerDTO.AnonymousDifferentiator)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if existingUser != nil {
			return errors.New("user already exists")
		}
	}

	hashedPassword, err := a.pwdService.HashPassword(registerDTO.Password)
	if err != nil {
		return err
	}
	new_user := &domain.User{
		ID:                      uuid.New(),
		FullName:                registerDTO.FullName,
		Email:                   registerDTO.Email,
		AnonymousDifferentiator: registerDTO.AnonymousDifferentiator,
		Password:                hashedPassword,
		PhoneNumber:             registerDTO.PhoneNumber,
		UserType:                registerDTO.UserType,
		Role:                    "regular",
		Active:                  true,
		Verified:                false,
		CounselorAssigned:       false,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	count, err := a.userRepo.GetUsersCount()
	if err != nil {
		return err
	}
	if count == 0 {
		new_user.Role = "admin"
		new_user.Verified = true
	}
	if existingUser != nil && existingUser.GoogleSignin {
		existingUser.FullName = new_user.FullName
		existingUser.Email = new_user.Email
		existingUser.Password = new_user.Password
		existingUser.PhoneNumber = new_user.PhoneNumber
		existingUser.UserType = new_user.UserType
		existingUser.Role = new_user.Role
		existingUser.Active = new_user.Active
		existingUser.Verified = new_user.Verified
		existingUser.CounselorAssigned = new_user.CounselorAssigned
		existingUser.UpdatedAt = time.Now()
		existingUser.GoogleSignin = false
		err = a.userRepo.UpdateUser(existingUser)
		if err != nil {
			return err
		}
		return nil
	}
	return a.userRepo.CreateUser(new_user)
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

	id := uuid.MustParse(claims["id"].(string))
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

	new_accessToken, new_refreshToken, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	user.AccessToken = new_accessToken
	user.RefreshToken = new_refreshToken
	err = a.userRepo.UpdateUser(user)
	if err != nil {
		return "", "", err
	}

	return new_accessToken, new_refreshToken, nil
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
		return errors.New("invalid token claims:code")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return errors.New("invalid token claims:email")
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
