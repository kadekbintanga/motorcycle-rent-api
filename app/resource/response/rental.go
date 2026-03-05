package response

import (
	"motorcycle-rent-api/app/model"
	"time"
)

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

type RentalListPaginationResponse struct {
	UUID                          string  `json:"uuid"`
	CustomerUUID                  string  `json:"customer_uuid"`
	MotocycleUUID                 string  `json:"motorcycle_uuid"`
	RentDate                      string  `json:"rent_date"`
	RentPrice                     float64 `json:"rent_price"`
	Status                        string  `json:"status"`
	CustomerNameCaptured          string  `json:"customer_name_captured"`
	CustomerIDNumberCaptured      string  `json:"customer_id_number_captured"`
	CustomerSIMNumberCaptured     string  `json:"customer_sim_number_captured"`
	CustomerPhoneCaptured         string  `json:"customer_phone_captured"`
	MotorcyclePlateNumberCaptured string  `json:"motorcycle_plate_number_captured"`
	CreatedAt                     string  `json:"created_at"`
	UpdatedAt                     string  `json:"updated_at"`
}

func RentalListPaginationFormatter(rentals []model.Rental) []RentalListPaginationResponse {
	var result []RentalListPaginationResponse
	for _, rental := range rentals {
		result = append(result, RentalListPaginationResponse{
			UUID:                          rental.UUID.String(),
			CustomerUUID:                  rental.CustomerUUID.String(),
			MotocycleUUID:                 rental.MotorcycleUUID.String(),
			RentDate:                      rental.RentDate.String(),
			RentPrice:                     rental.RentPrice,
			Status:                        string(rental.Status),
			CustomerNameCaptured:          rental.CustomerNameCaptured,
			CustomerIDNumberCaptured:      rental.CustomerIDNumberCaptured,
			CustomerSIMNumberCaptured:     rental.CustomerSIMNumberCaptured,
			CustomerPhoneCaptured:         rental.CustomerPhoneCaptured,
			MotorcyclePlateNumberCaptured: rental.MotorcyclePlateNumberCaptured,
			CreatedAt:                     rental.CreatedAt.String(),
			UpdatedAt:                     rental.UpdatedAt.String(),
		})
	}
	return result
}

type RentalPaymentResponse struct {
	UUID      string  `json:"uuid"`
	Amount    float64 `json:"amount"`
	Method    string  `json:"method"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type RentalDetailResponse struct {
	UUID                          string                  `json:"uuid"`
	CustomerUUID                  string                  `json:"customer_uuid"`
	MotocycleUUID                 string                  `json:"motorcycle_uuid"`
	RentDate                      string                  `json:"rent_date"`
	ReturnDatePlan                time.Time               `json:"return_date_plan"`
	ReturnDateActual              time.Time               `json:"return_date_actual"`
	LateDay                       int                     `json:"late_day"`
	PricePerDayCaptured           float64                 `json:"price_per_day_captured"`
	CustomerNameCaptured          string                  `json:"customer_name_captured"`
	CustomerIDNumberCaptured      string                  `json:"customer_id_number_captured"`
	CustomerSIMNumberCaptured     string                  `json:"customer_sim_number_captured"`
	CustomerPhoneCaptured         string                  `json:"customer_phone_captured"`
	MotorcyclePlateNumberCaptured string                  `json:"motorcycle_plate_number_captured"`
	RentPrice                     float64                 `json:"rent_price"`
	PenaltyPrice                  float64                 `json:"penalty_price"`
	Status                        string                  `json:"status"`
	Payment                       []RentalPaymentResponse `json:"payment"`
	CreatedAt                     time.Time               `json:"created_at"`
	UpdatedAt                     time.Time               `json:"updated_at"`
}

func RentalDetailFormatter(rental model.Rental) RentalDetailResponse {
	var payments []RentalPaymentResponse
	for _, payment := range rental.Payment {
		payments = append(payments, RentalPaymentResponse{
			UUID:      payment.UUID.String(),
			Amount:    payment.Amount,
			Method:    string(payment.Method),
			Type:      string(payment.Type),
			CreatedAt: payment.CreatedAt.String(),
			UpdatedAt: payment.UpdatedAt.String(),
		})
	}

	return RentalDetailResponse{
		UUID:                          rental.UUID.String(),
		CustomerUUID:                  rental.CustomerUUID.String(),
		MotocycleUUID:                 rental.MotorcycleUUID.String(),
		RentDate:                      rental.RentDate.String(),
		ReturnDatePlan:                rental.ReturnDatePlan,
		ReturnDateActual:              rental.ReturnDateActual,
		LateDay:                       rental.LateDay,
		PricePerDayCaptured:           rental.PricePerDayCaptured,
		CustomerNameCaptured:          rental.CustomerNameCaptured,
		CustomerIDNumberCaptured:      rental.CustomerIDNumberCaptured,
		CustomerSIMNumberCaptured:     rental.CustomerSIMNumberCaptured,
		CustomerPhoneCaptured:         rental.CustomerPhoneCaptured,
		MotorcyclePlateNumberCaptured: rental.MotorcyclePlateNumberCaptured,
		RentPrice:                     rental.RentPrice,
		PenaltyPrice:                  rental.PenaltyPrice,
		Status:                        string(rental.Status),
		Payment:                       payments,
		CreatedAt:                     rental.CreatedAt,
		UpdatedAt:                     rental.UpdatedAt,
	}
}

type RentalSummary struct {
	Total     int64 `json:"total"`
	Ongoing   int64 `json:"ongoing"`
	Completed int64 `json:"completed"`
	Canceled  int64 `json:"canceled"`
}
