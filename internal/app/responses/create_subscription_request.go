package responses

type CreateSubscriptionResponse struct {
	Status string `json:"status"`
	ID     int64  `json:"id"`
}

func ConvertToCreateSubscriptionResponse(ID int64) *CreateSubscriptionResponse {
	return &CreateSubscriptionResponse{Status: StatusOK, ID: ID}
}
