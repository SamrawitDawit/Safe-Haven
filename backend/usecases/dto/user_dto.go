package dto

import "errors"

type RegisterDTO struct {
	FullName                string `json:"fullName"`
	AnonymousDifferenitator string `json:"anonymousDifferenitator" binding:"required"`
	Email                   string `json:"email omitempty"`
	UserType                string `json:"userType" binding:"required"`
	Password                string `json:"password" binding:"required"`
	PhoneNumber             string `json:"phoneNumber"`
}
type LoginDTO struct {
	UserType                string `json:"userType" binding:"required"`
	Email                   string `json:"email" binding:"required"`
	AnonymousDifferenitator string `json:"anonymousDifferenitator" binding:"required"`
	PhoneNumber             string `json:"phoneNumber"`
	Password                string `json:"password" binding:"required"`
}

func ValidateRegisterDTO(registerDTO RegisterDTO) error {
	if registerDTO.UserType == "normal" {
		if registerDTO.FullName == "" {
			return errors.New("fullname is required")
		}
		if registerDTO.Email == "" && registerDTO.PhoneNumber == "" {
			return errors.New("email or phoneNumber is required")
		}
	} else if registerDTO.UserType == "anonymous" {
		if registerDTO.AnonymousDifferenitator == "" || registerDTO.Password == "" {
			return errors.New("differentiator and password is required")
		}
	}
	return nil
}

func ValidateLoginDTO(loginDTO LoginDTO) error {
	if loginDTO.UserType == "normal" {
		if loginDTO.Email == "" && loginDTO.PhoneNumber == "" {
			return errors.New("email or phoneNumber is required")
		}
		if loginDTO.Password == "" {
			return errors.New("password is required")
		}
	} else if loginDTO.UserType == "anonymous" {
		if loginDTO.AnonymousDifferenitator == "" || loginDTO.Password == "" {
			return errors.New("differentiator and password is required")
		}
	}
	return nil
}
