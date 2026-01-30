package responses

import "effective_mobile/internal/models"

type SubscriptionResponse struct {
	ID          int64  `json:"id"`
	ServiceName string `json:"service_name" `
	Price       int    `json:"price"`
	UserUUID    string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ConvertSubscription(sub *models.Subscription) *SubscriptionResponse {
	sr := new(SubscriptionResponse)
	sr.ID = sub.ID
	sr.ServiceName = sub.ServiceName
	sr.Price = sub.Price
	sr.UserUUID = sub.UserUUID
	sr.StartDate = sub.StartDate.Format("01-2006")
	if sub.EndDate != nil {
		sr.EndDate = sub.EndDate.Format("01-2006")
	} else {
		sr.EndDate = ""
	}
	sr.CreatedAt = sub.CreatedAt.Format("02-01-2006 15:04:05")
	sr.UpdatedAt = sub.UpdatedAt.Format("02-01-2006 15:04:05")
	return sr
}
