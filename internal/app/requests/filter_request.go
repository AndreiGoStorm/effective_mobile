package requests

import (
	"effective_mobile/internal/app/responses"
	"effective_mobile/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type FilterByUserAndServiceRequest struct {
	ServiceName string `query:"service_name" validate:"omitempty,alpha"`
	UserUUID    string `query:"user_id" validate:"omitempty,uuid"`
	StartDate   string `query:"start_date" validate:"required,datetime=01-2006"`
	EndDate     string `query:"end_date" validate:"required,datetime=01-2006"`
}

func (r *FilterByUserAndServiceRequest) GetFilter() *models.Filter {
	filter := &models.Filter{
		ServiceName: r.ServiceName,
		UserUUID:    r.UserUUID,
	}

	t, err := time.Parse("01-2006", r.StartDate)
	if err == nil {
		filter.StartDate = t
	}

	t, err = time.Parse("01-2006", r.EndDate)
	if err == nil {
		filter.EndDate = &t
	}

	return filter
}

func (r *FilterByUserAndServiceRequest) ValidationFilterError(errs validator.ValidationErrors) responses.Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "alpha":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not an alpha field", err.Field()))
		case "uuid":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a uuid field", err.Field()))
		case "datetime":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid datetime", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return responses.Response{Status: responses.StatusError, Error: strings.Join(errMsgs, ", ")}
}
