package constant

type CustomerStatus string

const (
	CustomerStatusActive      CustomerStatus = "ACTIVE"
	CustomerStatusBlacklisted CustomerStatus = "BLACKLISTED"
)
