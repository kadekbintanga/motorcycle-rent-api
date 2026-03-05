package constant

type PaymentType string

const (
	PaymentTypeDeposit     PaymentType = "DEPOSIT"
	PaymentTypeRentPayment PaymentType = "RENT_PAYMENT"
	PaymentTypeFullPayment PaymentType = "FULL_PAYMENT"
)

type PaymentMethod string

const (
	PaymentMethodCash     PaymentMethod = "CASH"
	PaymentMethodTransfer PaymentMethod = "TRANSFER"
	PaymentMethodQris     PaymentMethod = "QRIS"
)
