package requests

import (
	"effective_mobile/internal/app/responses"
	"effective_mobile/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateSubscriptionRequest struct {
	ServiceName string `json:"service_name" validate:"required,alpha" example:"YandexPlus"`
	Price       int    `json:"price" validate:"required,number,gt=0"  example:"400"`
	StartDate   string `json:"start_date" validate:"required,datetime=01-2006" example:"02-2026"`
	EndDate     string `json:"end_date" validate:"omitempty,datetime=01-2006" example:"04-2026"`
}

func (r *UpdateSubscriptionRequest) ValidationSubscriptionError(errs validator.ValidationErrors) responses.Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "alpha":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not an alpha field", err.Field()))
		case "datetime":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid datetime", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return responses.Error(strings.Join(errMsgs, ", "))
}

func (r *UpdateSubscriptionRequest) GetSubscription(sub *models.Subscription) *models.Subscription {
	sub.ServiceName = r.ServiceName
	sub.Price = r.Price

	t, err := time.Parse("01-2006", r.StartDate)
	if err == nil {
		sub.StartDate = t
	}

	sub.EndDate = nil
	t, err = time.Parse("01-2006", r.EndDate)
	if err == nil {
		sub.EndDate = &t
	}

	return sub
}
