package request

type CreateRentalRequest struct {
	CustomerUUID   string  `json:"customer_uuid" validate:"required,uuid"`
	MotorcycleUUID string  `json:"motorcycle_uuid" validate:"required,uuid"`
	Days           uint    `json:"days" validate:"required"`
	Payment        float64 `json:"payment" validate:"required"`
	PaymentType    string  `json:"payment_type" validate:"required,oneof=DEPOSIT RENT_PAYMENT"`
	PaymentMethod  string  `json:"payment_method" validate:"required,oneof=CASH TRANSFER QRIS"`
}

type ReturnRentalRequest struct {
	ReturnDate    string  `json:"return_date" validate:"required"`
	Payment       float64 `json:"payment"`
	PaymentMethod string  `json:"payment_method" validate:"payment_method_required,oneof=CASH TRANSFER QRIS"`
}

type RentalPaymentRequest struct {
	Payment       float64 `json:"payment" validate:"required"`
	PaymentMethod string  `json:"payment_method" validate:"oneof=CASH TRANSFER QRIS"`
}

type GetRentalListFilter struct {
	CustomerUUID   string `form:"customer_uuid" validate:"omitempty,uuid"`
	MotorcycleUUID string `form:"motorcycle_uuid" validate:"omitempty,uuid"`
	RentDateStart  string `form:"rent_date_start"`
	RentDateEnd    string `form:"rent_date_end"`
	Status         string `form:"status" validate:"omitempty,oneof=ONGOING COMPLETED CANCELED"`
}
