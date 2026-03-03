package response

import (
	"motorcycle-rent-api/app/model"
)

type FormattedAdminLogin struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatLoginResponse(user model.Admin, token string) FormattedAdminLogin {

	result := FormattedAdminLogin{
		UUID:  user.UUID.String(),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return result
}
