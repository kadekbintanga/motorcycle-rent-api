package constant

type MotorcycleStatus string

const (
	MotorcycleStatusAvailable   MotorcycleStatus = "AVAILABLE"
	MotorcycleStatusRented      MotorcycleStatus = "RENTED"
	MotorcycleStatusMaintenance MotorcycleStatus = "MAINTENANCE"
	MotorcycleStatusInactive    MotorcycleStatus = "INACTIVE"
)

type MotorcycleType string

const (
	MotorcycleTypeMatic  MotorcycleType = "MATIC"
	MotorcycleTypeManual MotorcycleType = "MANUAL"
)
