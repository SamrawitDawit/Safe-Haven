package usecases

import (
	"backend/domain"
	"backend/usecases/dto"
	"backend/usecases/interfaces"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AuthUseCaseInterface interface {
	Register(registerDTO dto.RegisterDTO) (*domain.User, *domain.CustomError)
	Login(loginDTO dto.LoginDTO) (*domain.User, string, string, *domain.CustomError)
	RefreshToken(refreshToken string) (string, string, *domain.CustomError)
	ForgotPassword(email string) *domain.CustomError
	ResetPassword(token string, newPassword string) *domain.CustomError
	HandleGoogleCallback(user *domain.User) (string, string, *domain.CustomError)
}

func NewAuthUseCase(userRepo interfaces.UserRepositoryInterface, jwtservice interfaces.JwtServiceInterface, emailService interfaces.EmailServiceInterface, pwdService interfaces.HashingServiceInterface, encryptionService interfaces.EncryptionServiceInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo:          userRepo,
		jwtService:        jwtservice,
		emailService:      emailService,
		pwdService:        pwdService,
		encryptionService: encryptionService,
	}
}

type AuthUseCase struct {
	userRepo          interfaces.UserRepositoryInterface
	jwtService        interfaces.JwtServiceInterface
	emailService      interfaces.EmailServiceInterface
	pwdService        interfaces.HashingServiceInterface
	encryptionService interfaces.EncryptionServiceInterface
}

func (a *AuthUseCase) validateRegisterDTO(registerDTO dto.RegisterDTO) *domain.CustomError {
	if registerDTO.Email == "" && registerDTO.PhoneNumber == "" {
		return domain.ErrEmailOrPhoneRequired
	}
	return nil
}

func (a *AuthUseCase) validateLoginDTO(loginDTO dto.LoginDTO) *domain.CustomError {
	if loginDTO.Email == "" && loginDTO.PhoneNumber == "" {
		return domain.ErrEmailOrPhoneRequired
	}
	return nil
}

// Password validation function
func validatePassword(password string) *domain.CustomError {
	// Minimum 8 characters, at least one uppercase letter, one lowercase letter, one number, and one special character
	var (
		minLength        = 8
		uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
		lowercaseRegex   = regexp.MustCompile(`[a-z]`)
		numberRegex      = regexp.MustCompile(`[0-9]`)
		specialCharRegex = regexp.MustCompile(`[!@#~$%^&*()+|_.,]`)
	)

	if len(password) < minLength {
		return domain.NewCustomError("password must be at least 8 characters long", http.StatusBadRequest)
	}
	if !uppercaseRegex.MatchString(password) {
		return domain.NewCustomError("password must contain at least one uppercase letter", http.StatusBadRequest)
	}
	if !lowercaseRegex.MatchString(password) {
		return domain.NewCustomError("password must contain at least one lowercase letter", http.StatusBadRequest)
	}
	if !numberRegex.MatchString(password) {
		return domain.NewCustomError("password must contain at least one number", http.StatusBadRequest)
	}
	if !specialCharRegex.MatchString(password) {
		return domain.NewCustomError("password must contain at least one special character", http.StatusBadRequest)
	}

	return nil
}

func (a *AuthUseCase) generateAndStoreTokens(user *domain.User) (string, string, *domain.CustomError) {
	accessToken, refreshToken, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	encryptedAccessToken, err := a.encryptionService.Encrypt(accessToken)
	if err != nil {
		return "", "", err
	}
	encryptedRefreshToken, err := a.encryptionService.Encrypt(refreshToken)
	if err != nil {
		return "", "", err
	}

	user.AccessToken = encryptedAccessToken
	user.RefreshToken = encryptedRefreshToken

	updated_fields := map[string]interface{}{
		"accessToken":  encryptedAccessToken,
		"refreshToken": encryptedRefreshToken,
	}
	err = a.userRepo.UpdateUserFields(user.ID, updated_fields)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthUseCase) generateResetToken(email string) (int64, string, *domain.CustomError) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min, max := 10000, 100000
	randomNumber := r.Int63n(int64(max-min)) + int64(min)
	resetToken, err := a.jwtService.GenerateResetToken(email, randomNumber)
	if err != nil {
		return 0, "", err
	}
	return randomNumber, resetToken, nil
}

func (a *AuthUseCase) Login(loginDTO dto.LoginDTO) (*domain.User, string, string, *domain.CustomError) {
	var user *domain.User
	var err *domain.CustomError
	err = a.validateLoginDTO(loginDTO)
	if err != nil {
		return nil, "", "", err
	}
	if loginDTO.Email != "" {
		user, err = a.userRepo.GetUserByEmail(loginDTO.Email)
		if err != nil {
			return nil, "", "", err
		}
		if user != nil && user.GoogleSignin {
			return nil, "", "", domain.ErrInvalidCredentials
		}
	} else if loginDTO.PhoneNumber != "" {
		user, err = a.userRepo.GetUserByPhoneNumber(loginDTO.PhoneNumber)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = a.pwdService.CheckHash(loginDTO.Password, user.Password)
	if err != nil {
		return nil, "", "", err
	}

	accessToken, refreshToken, err := a.generateAndStoreTokens(user)
	if err != nil {
		return nil, "", "", err
	}
	return user, accessToken, refreshToken, nil
}

func (a *AuthUseCase) Register(registerDTO dto.RegisterDTO) (*domain.User, *domain.CustomError) {
	var existingUser *domain.User
	var err *domain.CustomError
	//checking if user already exists
	err = a.validateRegisterDTO(registerDTO)
	if err != nil {
		return nil, err
	}
	// Password validation
	err = validatePassword(registerDTO.Password)
	if err != nil {
		return nil, err
	}
	if registerDTO.Email != "" {
		existingUser, err = a.userRepo.GetUserByEmail(registerDTO.Email)
		if err != nil && err != domain.ErrUserNotFound {
			return nil, err
		}
		if existingUser != nil && !existingUser.GoogleSignin {
			return nil, domain.ErrUserEmailExists
		}

	} else if registerDTO.PhoneNumber != "" {
		existingUser, err = a.userRepo.GetUserByPhoneNumber(registerDTO.PhoneNumber)
		if err != nil && err != domain.ErrUserNotFound {
			return nil, err
		}
		if existingUser != nil {
			return nil, domain.ErrUserPhoneNumberExists
		}
	}

	hashedPassword, err := a.pwdService.Hash(registerDTO.Password)
	if err != nil {
		return nil, err
	}
	new_user := &domain.User{
		ID:                uuid.New(),
		FullName:          registerDTO.FullName,
		Email:             registerDTO.Email,
		Password:          hashedPassword,
		PhoneNumber:       registerDTO.PhoneNumber,
		Category:          registerDTO.Category,
		Language:          registerDTO.Language,
		Role:              "regular",
		Active:            true,
		Verified:          false,
		CounselorAssigned: false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	count, err := a.userRepo.GetUsersCount()
	if err != nil {
		return nil, err
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
		existingUser.Role = new_user.Role
		existingUser.Active = new_user.Active
		existingUser.Verified = new_user.Verified
		existingUser.CounselorAssigned = new_user.CounselorAssigned
		existingUser.UpdatedAt = time.Now()
		existingUser.GoogleSignin = false
		err = a.userRepo.UpdateUser(existingUser)
		if err != nil {
			return nil, err
		}
		return existingUser, nil
	}
	err = a.userRepo.CreateUser(new_user)
	if err != nil {
		return nil, err
	}
	return new_user, nil
}

func (a *AuthUseCase) RefreshToken(refreshToken string) (string, string, *domain.CustomError) {
	token, err := a.jwtService.ValidateToken(refreshToken)
	if err != nil || !token.Valid {
		return "", "", domain.ErrInvalidToken
	}
	claims, err := a.jwtService.ExtractTokenClaims(token)
	if err != nil {
		return "", "", err
	}

	id := uuid.MustParse(claims["id"].(string))
	user, err := a.userRepo.GetUserByIDWithLock(id)
	if err != nil || user == nil {
		return "", "", err
	}
	defer a.releaseUserLock(user.ID)
	token_from_DB := user.RefreshToken
	decryptedToken, err := a.encryptionService.Decrypt(token_from_DB)
	if err != nil {
		return "", "", err
	}
	if decryptedToken != refreshToken {
		return "", "", domain.ErrInvalidRefreshToken
	}

	accessToken, refreshToken, err := a.generateAndStoreTokens(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
func (a *AuthUseCase) releaseUserLock(userID uuid.UUID) *domain.CustomError {
	updated_fields := map[string]interface{}{
		"lock": false,
	}
	err := a.userRepo.UpdateUserFields(userID, updated_fields)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUseCase) ForgotPassword(email string) *domain.CustomError {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.ResetToken != "" && user.ResetTokenExpiry.After(time.Now()) {
		return domain.ErrResetTokenAlreadySent
	}
	// send email
	randomNumber, resetToken, err := a.generateResetToken(email)
	if err != nil {
		return err
	}
	encryptedResetToken, err := a.encryptionService.Encrypt(resetToken)
	if err != nil {
		return err
	}
	code := strconv.FormatInt(randomNumber, 10)
	hashedCode, err := a.pwdService.Hash(code)
	if err != nil {
		return err
	}
	expiryTime := time.Now().Add(time.Minute * 5)
	updated_fields := map[string]interface{}{
		"resetCode":        hashedCode,
		"resetToken":       encryptedResetToken,
		"resetTokenExpiry": expiryTime,
	}

	err = a.userRepo.UpdateUserFields(user.ID, updated_fields)
	if err != nil {
		return err
	}
	return a.emailService.SendResetPasswordEmail(user.Email, resetToken)
}

func (a *AuthUseCase) ResetPassword(token string, newPassword string) *domain.CustomError {
	validatedToken, err := a.jwtService.ValidateToken(token)
	if err != nil || !validatedToken.Valid {
		return domain.ErrInvalidToken
	}
	claims, err := a.jwtService.ExtractTokenClaims(validatedToken)
	if err != nil {
		return err
	}
	code, ok := claims["code"].(float64)
	if !ok {
		return domain.NewCustomError("Invalid code", http.StatusBadRequest)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return domain.NewCustomError("Invalid email", http.StatusBadRequest)
	}
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return err
	}
	codestr := strconv.FormatInt(int64(code), 10)
	err = a.pwdService.CheckHash(codestr, user.ResetCode)
	if err != nil {
		return domain.ErrInvalidResetCode
	}
	stored_reset_token := user.ResetToken
	decryptedToken, err := a.encryptionService.Decrypt(stored_reset_token)
	if err != nil {
		return err
	}
	if decryptedToken != token {
		return domain.ErrInvalidResetToken
	}
	//validating password
	err = validatePassword(newPassword)
	if err != nil {
		return err
	}
	hashedPassword, err := a.pwdService.Hash(newPassword)
	if err != nil {
		return err
	}

	updated_fields := map[string]interface{}{
		"password":         string(hashedPassword),
		"resetCode":        "",
		"resetToken":       "",
		"resetTokenExpiry": nil,
	}

	err = a.userRepo.UpdateUserFields(user.ID, updated_fields)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUseCase) HandleGoogleCallback(user *domain.User) (string, string, *domain.CustomError) {
	existingUser, err := a.userRepo.GetUserByEmail(user.Email)
	if err != nil && err != domain.ErrUserNotFound {
		return "", "", err
	}
	if existingUser != nil {
		if !user.GoogleSignin {
			return "", "", domain.ErrUserEmailExists
		}
		accessToken, refreshToken, err := a.generateAndStoreTokens(existingUser)
		if err != nil {
			return "", "", err
		}
		return accessToken, refreshToken, nil

	}
	accessToken, refreshToken, err := a.generateAndStoreTokens(user)
	if err != nil {
		return "", "", err
	}

	user.ID = uuid.New()
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
