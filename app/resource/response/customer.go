package response

import "motorcycle-rent-api/app/model"

type CustomerListPaginationResponse struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	IDNumber  string `json:"id_number"`
	SIMNumber string `json:"sim_number"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CustomerListPaginationFormatter(customers []model.Customer) []CustomerListPaginationResponse {
	var result []CustomerListPaginationResponse
	for _, customer := range customers {
		result = append(result, CustomerListPaginationResponse{
			UUID:      customer.UUID.String(),
			Name:      customer.Name,
			IDNumber:  customer.IDNumber,
			SIMNumber: customer.SIMNumber,
			Phone:     customer.Phone,
			Status:    string(customer.Status),
			CreatedAt: customer.CreatedAt.String(),
			UpdatedAt: customer.UpdatedAt.String(),
		})
	}
	return result
}

type CustomerDetailResponse struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	IDNumber        string `json:"id_number"`
	SIMNumber       string `json:"sim_number"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	Status          string `json:"status"`
	BlacklistReason string `json:"blacklist_reason"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func CustomerDetailFormatter(customer model.Customer) CustomerDetailResponse {
	return CustomerDetailResponse{
		UUID:            customer.UUID.String(),
		Name:            customer.Name,
		IDNumber:        customer.IDNumber,
		SIMNumber:       customer.SIMNumber,
		Phone:           customer.Phone,
		Address:         customer.Address,
		Status:          string(customer.Status),
		BlacklistReason: customer.BlacklistReason,
		CreatedAt:       customer.CreatedAt.String(),
		UpdatedAt:       customer.UpdatedAt.String(),
	}
}

type CustomerSummary struct {
	Total       int64 `json:"total"`
	Active      int64 `json:"active"`
	Blacklisted int64 `json:"blacklisted"`
}
