package infrastructure

import (
	"backend/delivery/config"
	"context"
	"fmt"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchapb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
)

type RecaptchaService struct{}

func (r *RecaptchaService) CreateAssessment(token string) (float32, error) {
	projectID := config.ENV.PROJECT_ID
	recaptchaKey := config.ENV.RECAPTCHA_KEY
	recaptchaAction := config.ENV.RECAPTCHA_ACTION
	ctx := context.Background()
	client, err := recaptcha.NewClient(ctx)
	if err != nil {
		return 0, fmt.Errorf("error creating reCAPTCHA client: %v", err)
	}
	defer client.Close()

	event := &recaptchapb.Event{
		Token:   token,
		SiteKey: recaptchaKey,
	}

	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: &recaptchapb.Assessment{Event: event},
		Parent:     fmt.Sprintf("projects/%s", projectID),
	}

	response, err := client.CreateAssessment(ctx, request)
	if err != nil {
		return 0, fmt.Errorf("error calling CreateAssessment: %v", err)
	}

	if !response.TokenProperties.Valid {
		return 0, fmt.Errorf("invalid token: %v", response.TokenProperties.InvalidReason)
	}

	if response.TokenProperties.Action != recaptchaAction {
		return 0, fmt.Errorf("unexpected action: %s", response.TokenProperties.Action)
	}

	score := response.RiskAnalysis.Score
	return score, nil
}
