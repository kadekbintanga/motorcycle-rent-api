package response

import "motorcycle-rent-api/app/model"

type MotorcycleListpaginationResponse struct {
	UUID        string `json:"uuid"`
	PlateNumber string `json:"plate_number"`
	Brand       string `json:"brand"`
	Type        string `json:"type"`
	Year        int    `json:"year"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func MotorcycleListpaginationResponseFormatter(motorcycles []model.Motorcycle) []MotorcycleListpaginationResponse {
	var result []MotorcycleListpaginationResponse
	for _, motormotorcycle := range motorcycles {
		result = append(result, MotorcycleListpaginationResponse{
			UUID:        motormotorcycle.UUID.String(),
			PlateNumber: motormotorcycle.PlateNumber,
			Brand:       motormotorcycle.Brand,
			Type:        string(motormotorcycle.Type),
			Year:        motormotorcycle.Year,
			Status:      string(motormotorcycle.Status),
			CreatedAt:   motormotorcycle.CreatedAt.String(),
			UpdatedAt:   motormotorcycle.UpdatedAt.String(),
		})
	}
	return result
}
