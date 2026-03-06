package request

type CreateCustomerRequest struct {
	Name      string `json:"name" validate:"required,not_only_space,min=2,max=100,name_validation"`
	IDNumber  string `json:"id_number" validate:"required,not_only_space,min=6,max=20,alphanum"`
	SIMNumber string `json:"sim_number" validate:"required,not_only_space,min=5,max=20,alphanum"`
	Phone     string `json:"phone" validate:"required,not_only_space,min=10,max=15,phone_validation"`
	Address   string `json:"address" validate:"required,not_only_space,min=10,max=200"`
}

type UpdateCustomerDetailRequest struct {
	Name            string `json:"name" validate:"required,not_only_space,min=2,max=100,name_validation"`
	IDNumber        string `json:"id_number" validate:"required,not_only_space,min=6,max=20,alphanum"`
	SIMNumber       string `json:"sim_number" validate:"required,not_only_space,min=5,max=20,alphanum"`
	Phone           string `json:"phone" validate:"required,not_only_space,min=10,max=15,phone_validation"`
	Address         string `json:"address" validate:"required,not_only_space,min=10,max=200"`
	Status          string `json:"status" validate:"oneof=ACTIVE BLACKLISTED"`
	BlacklistReason string `json:"blacklist_reason" validate:"required_if=Status BLACKLISTED"`
}

type UpdateCustomerStatusRequest struct {
	Status          string `json:"status" validate:"oneof=ACTIVE BLACKLISTED"`
	BlacklistReason string `json:"blacklist_reason" validate:"required_if=Status BLACKLISTED"`
}

type GetCustomerListFilter struct {
	Status string `form:"status" validate:"omitempty,oneof=ACTIVE BLACKLISTED"`
}
