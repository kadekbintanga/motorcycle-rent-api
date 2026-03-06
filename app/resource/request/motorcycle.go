package request

type CreateMotorcycleRequest struct {
	PlateNumber string  `json:"plate_number" validate:"required,plate_number"`
	Brand       string  `json:"brand" validate:"required,not_only_space,min=1,max=100"`
	Type        string  `json:"type" validate:"required,oneof=MANUAL MATIC"`
	Year        int     `json:"year" validate:"required"`
	Status      string  `json:"status" validate:"required,oneof=AVAILABLE MAINTENANCE INACTIVE"`
	PricePerDay float64 `json:"price_per_day" validate:"required"`
}

type GetMotorcycleListFilter struct {
	Type   string `form:"type" validate:"omitempty,oneof=MANUAL MATIC"`
	Status string `form:"status" validate:"omitempty,oneof=AVAILABLE MAINTENANCE INACTIVE RENTED"`
	Year   int    `form:"year" validate:"omitempty"`
}

type UpdateMotorcycleRequest struct {
	PlateNumber string  `json:"plate_number" validate:"required,not_only_space,plate_number"`
	Brand       string  `json:"brand" validate:"required,not_only_space,min=1,max=100"`
	Type        string  `json:"type" validate:"oneof=MANUAL MATIC"`
	Year        int     `json:"year" validate:"required"`
	Status      string  `json:"status" validate:"oneof=AVAILABLE MAINTENANCE INACTIVE"`
	PricePerDay float64 `json:"price_per_day" validate:"required"`
}

type UpdateMotorcycleStatusRequest struct {
	Status string `json:"status" validate:"oneof=AVAILABLE MAINTENANCE INACTIVE"`
}
