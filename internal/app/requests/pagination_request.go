package requests

import (
	"effective_mobile/internal/app/responses"
	"effective_mobile/internal/models"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type PaginationRequest struct {
	Page int `json:"page" validate:"omitempty,number,min=1"`
	Size int `json:"size" validate:"omitempty,number,min=1"`
}

func (r *PaginationRequest) setDefault() {
	if r.Page < 1 {
		r.Page = 1
	}
	validSizes := map[int]bool{5: true, 10: true, 15: true, 20: true}
	if !validSizes[r.Size] {
		r.Size = 5
	}
}

func (r *PaginationRequest) GetSubscriptionPagination() *models.SubscriptionPagination {
	r.setDefault()
	return &models.SubscriptionPagination{
		Page: r.Page,
		Size: r.Size,
	}
}

func (r *PaginationRequest) ValidationPaginationError(errs validator.ValidationErrors) responses.Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "number":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a number field", err.Field()))
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is less than min", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return responses.Error(strings.Join(errMsgs, ", "))
}
