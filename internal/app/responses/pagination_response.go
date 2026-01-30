package responses

import (
	"effective_mobile/internal/models"
)

type PaginatedResponse struct {
	Page  int                     `json:"page"`
	Size  int                     `json:"size"`
	Total int                     `json:"total"`
	Data  []*SubscriptionResponse `json:"data"`
}

func NewPaginatedResponse(page, size int) *PaginatedResponse {
	return &PaginatedResponse{Page: page, Size: size}
}

func (p *PaginatedResponse) ConvertSubscriptionsToResponse(total int, subs []*models.Subscription) *PaginatedResponse {
	p.Total = total
	data := make([]*SubscriptionResponse, 0)
	for _, subscription := range subs {
		data = append(data, ConvertSubscription(subscription))
	}
	p.Data = data
	return p
}
