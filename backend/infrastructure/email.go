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

func (s *EmailService) BuildResetPasswordEmail(resetToken string) string {
	resetLink := "http://localhost:8080/reset-password?token=" + resetToken

	return `
		<!DOCTYPE html>
		<html>
		    <head></head>
		    <body>
		        <div>
			        <h1>Reset Password</h1>
				    <p>Hello,</p>
				    <p>We received a request to reset your password. Click the button below to reset it:</p>
				    <a href="` + resetLink + `">Reset Password</a>
				    <p>If you did not request this, please ignore this email.</p>
				    <p>Thanks,</p>
			    </div>
		    </body>
	    </html>`
}

func (s *EmailService) SendResetPasswordEmail(email string, resetToken string) *domain.CustomError {
	body := s.BuildResetPasswordEmail(resetToken)

	from := s.Email
	to := email

	subject := "Reset Password"
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
