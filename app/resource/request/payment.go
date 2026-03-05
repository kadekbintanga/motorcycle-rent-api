package request

type GetPaymentListFilter struct {
	CustomerUUID   string `form:"customer_uuid" validate:"omitempty,uuid"`
	MotorcycleUUID string `form:"motorcycle_uuid" validate:"omitempty,uuid"`
	DateStart      string `form:"date_start"`
	DateEnd        string `form:"date_end"`
	Method         string `form:"method" validate:"omitempty,oneof=CASH TRANSFER QRIS"`
	Type           string `form:"type" validate:"omitempty,oneof=DEPOSIT RENT_PAYMENT FULL_PAYMENT"`
}

type GetPaymentSummaryFilter struct {
	DateStart string `form:"date_start"`
	DateEnd   string `form:"date_end"`
}
