package responses

type FilterResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

func ConvertToFilterResponse(total int) *FilterResponse {
	return &FilterResponse{Status: StatusOK, Total: total}
}
