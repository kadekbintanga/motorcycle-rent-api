package service

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"

	"gorm.io/gorm"
)

type AdminServiceInterface interface {
	Login(apiCallID string, payload request.AdminLoginRequest) (*response.FormattedAdminLogin, constant.ResponseMap)
}

type AdminService struct {
	DB              *gorm.DB
	Config          *global.EnvConfig
	AdminRepository repository.AdminRepositoryInterface
}

func NewAdminService(db *gorm.DB, config *global.EnvConfig, adminRepository repository.AdminRepositoryInterface) AdminServiceInterface {
	return &AdminService{
		DB:              db,
		Config:          config,
		AdminRepository: adminRepository,
	}
}

func (a *AdminService) Login(apiCallID string, payload request.AdminLoginRequest) (*response.FormattedAdminLogin, constant.ResponseMap) {
	adminFound, err := a.AdminRepository.GetAdminByEmail(a.DB, payload.Email)
	if err != nil || adminFound == nil {
		helper.LogError(apiCallID, "Admin not found with email: "+payload.Email)
		return nil, constant.Res400InvalidEmailorPassword
	}

	err = helper.VerifyPassword(adminFound.Password, payload.Password)
	if err != nil {
		helper.LogError(apiCallID, "Invalid password for email: "+payload.Email)
		return nil, constant.Res400InvalidEmailorPassword
	}

	token, err := helper.GenerateJWTAdmin(adminFound, a.Config.JWTSecretAdmin, a.Config.JWTExpiredDurationAdmin)
	if err != nil {
		helper.LogError(apiCallID, "Failed to generate JWT token for email: "+payload.Email+" with error: "+err.Error())
		return nil, constant.Res422SomethingWentWrong
	}

	formattedResponse := response.FormatLoginResponse(*adminFound, token)

	return &formattedResponse, constant.Res200Success
}
