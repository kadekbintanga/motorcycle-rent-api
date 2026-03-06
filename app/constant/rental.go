package constant

type RentalStatus string

const (
	RentalStatusOngoing   RentalStatus = "ONGOING"
	RentalStatusCompleted RentalStatus = "COMPLETED"
	RentalStatusCancelled RentalStatus = "CANCELED"
)

type RetalTablePreload string

const (
	RentalPayment RetalTablePreload = "Payment"
)
