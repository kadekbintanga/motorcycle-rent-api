package request

type CreateMotorcycleRequest struct {
	PlateNumber string `json:"plate_number" validate:"required,plate_number"`
	Brand       string `json:"brand" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=MANUAL MATIC"`
	Year        int    `json:"year" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=AVAILABLE MAINTENANCE INACTIVE"`
}
