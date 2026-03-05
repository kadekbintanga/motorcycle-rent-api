package request

type CreateRentalRequest struct {
	CustomerUUID   string  `json:"customer_uuid" validate:"required,uuid"`
	MotorcycleUUID string  `json:"motorcycle_uuid" validate:"required,uuid"`
	Days           uint    `json:"days" validate:"required"`
	Payment        float64 `json:"payment" validate:"required"`
	PaymentType    string  `json:"payment_type" validate:"required,oneof=DEPOSIT RENT_PAYMENT"`
	PaymentMethod  string  `json:"payment_method" validate:"required,oneof=CASH TRANSFER QRIS"`
}
