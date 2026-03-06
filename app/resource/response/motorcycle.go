package response

import "motorcycle-rent-api/app/model"

type MotorcycleListpaginationResponse struct {
	UUID        string  `json:"uuid"`
	PlateNumber string  `json:"plate_number"`
	Brand       string  `json:"brand"`
	Type        string  `json:"type"`
	Year        int     `json:"year"`
	Status      string  `json:"status"`
	PricePerDay float64 `json:"price_per_day"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func MotorcycleListpaginationResponseFormatter(motorcycles []model.Motorcycle) []MotorcycleListpaginationResponse {
	var result []MotorcycleListpaginationResponse
	for _, motorcycle := range motorcycles {
		result = append(result, MotorcycleListpaginationResponse{
			UUID:        motorcycle.UUID.String(),
			PlateNumber: motorcycle.PlateNumber,
			Brand:       motorcycle.Brand,
			Type:        string(motorcycle.Type),
			Year:        motorcycle.Year,
			Status:      string(motorcycle.Status),
			PricePerDay: motorcycle.PricePerDay,
			CreatedAt:   motorcycle.CreatedAt.String(),
			UpdatedAt:   motorcycle.UpdatedAt.String(),
		})
	}
	return result
}

type MotorcycleSummary struct {
	Total       int `json:"total"`
	Available   int `json:"available"`
	Rented      int `json:"rented"`
	Maintenance int `json:"maintenance"`
	Inactive    int `json:"inactive"`
}
