package interfaces

type RecaptchaInterface interface {
	CreateAssessment(token string) (float32, error)
}
