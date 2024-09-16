package infrastructure

import (
	"backend/domain"
	"net/smtp"
)

type EmailService struct {
	Email    string
	UserName string
	Password string
	Host     string
	Port     string
}

func (s *EmailService) SendResetPasswordEmail(email string, resetToken string) *domain.CustomError {
	resetLink := "http://localhost:8080/reset-password?token=" + resetToken

	from := s.Email
	to := email

	subject := "Reset Password"
	body := `
		<!DOCTYPE html>
		<html>
		    <head>
		        <style>
		        </style>

		    </head>
		    <body>
		        <div class = "container">
			        <h1>Reset Password</h1>
				    <p>Hello</p>
				    <p> We received a request to reset your password. Click the button below to reset it: </p>
				    <a href="` + resetLink + `">Reset Password</a>
				    <p>If you did not request a password reset, please ignore this email.</p>
				    <p>Thanks</p>
			    </div>
		    </body>
	    </html>
				`

	//MIME Headers
	message := "MIME-Version: 1.0" + "\r\n"
	message += "Content-type: text/html; charset=\"UTF-8\"\r\n"
	message += "From: " + from + "\r\n"
	message += "To: " + to + "\r\n"
	message += "Subject: " + subject + "\r\n\r\n"
	message += body

	auth := smtp.PlainAuth("", s.UserName, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, from, []string{to}, []byte(message))
	if err != nil {
		return domain.ErrEmailSendingFailed
	}
	return nil
}
