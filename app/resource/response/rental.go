package response

import "motorcycle-rent-api/app/model"

type CreateRentalResponse struct {
	UUID                          string  `json:"uuid"`
	CustomerUUID                  string  `json:"customer_uuid"`
	MotocycleUUID                 string  `json:"motorcycle_uuid"`
	RentDate                      string  `json:"rent_date"`
	ReturnDatePlan                string  `json:"rent_date_plan"`
	RentPrice                     float64 `json:"rent_price"`
	Status                        string  `json:"status"`
	PricePerDayCaptured           float64 `json:"price_per_day_captured"`
	CustomerNameCaptured          string  `json:"customer_name_captured"`
	CustomerIDNumberCaptured      string  `json:"customer_id_number_captured"`
	CustomerSIMNumberCaptured     string  `json:"customer_sim_number_captured"`
	CustomerPhoneCaptured         string  `json:"customer_phone_captured"`
	MotorcyclePlateNumberCaptured string  `json:"motorcycle_plate_number_captured"`
	CreatedAt                     string  `json:"created_at"`
	UpdatedAt                     string  `json:"updated_at"`
}

func CreateRentalResponseFormatter(rental model.Rental) CreateRentalResponse {
	return CreateRentalResponse{
		UUID:                          rental.UUID.String(),
		CustomerUUID:                  rental.CustomerUUID.String(),
		MotocycleUUID:                 rental.MotorcycleUUID.String(),
		RentDate:                      rental.RentDate.String(),
		ReturnDatePlan:                rental.ReturnDatePlan.String(),
		RentPrice:                     rental.RentPrice,
		Status:                        string(rental.Status),
		PricePerDayCaptured:           rental.PricePerDayCaptured,
		CustomerNameCaptured:          rental.CustomerNameCaptured,
		CustomerIDNumberCaptured:      rental.CustomerIDNumberCaptured,
		CustomerSIMNumberCaptured:     rental.CustomerSIMNumberCaptured,
		CustomerPhoneCaptured:         rental.CustomerPhoneCaptured,
		MotorcyclePlateNumberCaptured: rental.MotorcyclePlateNumberCaptured,
		CreatedAt:                     rental.CreatedAt.String(),
		UpdatedAt:                     rental.UpdatedAt.String(),
	}
}
