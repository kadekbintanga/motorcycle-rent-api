package response

import (
	"motorcycle-rent-api/app/model"
	"time"
)

type PaymentListResponse struct {
	UUID      string    `json:"uuid"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PaymentListPaginationFormatter(payments []model.Payment) []PaymentListResponse {
	var result []PaymentListResponse
	for _, payment := range payments {
		result = append(result, PaymentListResponse{
			UUID:      payment.UUID.String(),
			Amount:    payment.Amount,
			Method:    string(payment.Method),
			Type:      string(payment.Type),
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		})
	}
	return result
}

type PaymentRentalResponse struct {
	UUID                          string                  `json:"uuid"`
	CustomerUUID                  string                  `json:"customer_uuid"`
	MotocycleUUID                 string                  `json:"motorcycle_uuid"`
	RentDate                      time.Time               `json:"rent_date"`
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

type PaymentDetailtResponse struct {
	UUID      string                `json:"uuid"`
	Amount    float64               `json:"amount"`
	Method    string                `json:"method"`
	Type      string                `json:"type"`
	Rental    PaymentRentalResponse `json:"rental"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func PaymentDetailFormatter(payment model.Payment) PaymentDetailtResponse {
	rental := PaymentRentalResponse{
		UUID:                          payment.Rental.UUID.String(),
		CustomerUUID:                  payment.Rental.CustomerUUID.String(),
		MotocycleUUID:                 payment.Rental.MotorcycleUUID.String(),
		RentDate:                      payment.Rental.RentDate,
		ReturnDatePlan:                payment.Rental.ReturnDatePlan,
		ReturnDateActual:              payment.Rental.ReturnDateActual,
		LateDay:                       payment.Rental.LateDay,
		PricePerDayCaptured:           payment.Rental.PricePerDayCaptured,
		CustomerNameCaptured:          payment.Rental.CustomerNameCaptured,
		CustomerIDNumberCaptured:      payment.Rental.CustomerIDNumberCaptured,
		CustomerSIMNumberCaptured:     payment.Rental.CustomerSIMNumberCaptured,
		CustomerPhoneCaptured:         payment.Rental.CustomerPhoneCaptured,
		MotorcyclePlateNumberCaptured: payment.Rental.MotorcyclePlateNumberCaptured,
		RentPrice:                     payment.Rental.RentPrice,
		PenaltyPrice:                  payment.Rental.PenaltyPrice,
		Status:                        string(payment.Rental.Status),
		CreatedAt:                     payment.Rental.CreatedAt,
		UpdatedAt:                     payment.Rental.UpdatedAt,
	}
	return PaymentDetailtResponse{
		UUID:      payment.UUID.String(),
		Amount:    payment.Amount,
		Method:    string(payment.Method),
		Type:      string(payment.Type),
		Rental:    rental,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}

type PaymentSummary struct {
	TotalAmount   float64 `json:"total_amount"`
	TotalCash     float64 `json:"total_cash"`
	TotalTransfer float64 `json:"total_transfer"`
	TotalQris     float64 `json:"total_qris"`
}
