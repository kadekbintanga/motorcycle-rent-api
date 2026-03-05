package constant

type RentalStatus string

const (
	RentalStatusOngoing   RentalStatus = "ONGOING"
	RentalStatusCompleted RentalStatus = "COMPLETED"
	RentalStatusCancelled RentalStatus = "CANCELLED"
)
