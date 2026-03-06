package request

type AdminLoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,cms_admin_password"`
}
