package service

import (
	"errors"
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerServiceInterface interface {
	CreateCustomer(apiCallID string, payload request.CreateCustomerRequest) constant.ResponseMap
	UpdateCustomerDetail(apiCallID string, customerUUID string, payload request.UpdateCustomerDetailRequest) constant.ResponseMap
	UpdateCustomerStatus(apiCallID string, customerUUID string, payload request.UpdateCustomerStatusRequest) constant.ResponseMap
	GetListCustomerPagination(apiCallID string, param helper.PaginationParam, filter request.GetCustomerListFilter) ([]response.CustomerListPaginationResponse, *helper.ResponseMeta, constant.ResponseMap)
	GetDetailCustomer(apiCallID string, customerUUID string) (*response.CustomerDetailResponse, constant.ResponseMap)
}

type CustomerService struct {
	DB                 *gorm.DB
	CustomerRepository repository.CustomerRepositoryInterface
}

func NewCustomerService(db *gorm.DB, customerRepository repository.CustomerRepositoryInterface) CustomerServiceInterface {
	return &CustomerService{
		DB:                 db,
		CustomerRepository: customerRepository,
	}
}

func (c *CustomerService) CreateCustomer(apiCallID string, payload request.CreateCustomerRequest) constant.ResponseMap {
	err := c.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		checkIDSIMNumber, err := c.CustomerRepository.GetCustomerByIDNumberORSIMNumber(tx, payload.IDNumber, payload.SIMNumber, "")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogError(apiCallID, "Error check ID Number and SIM Number: "+err.Error())
			return errors.New("error get check ID and SIM Number")
		}

		if checkIDSIMNumber != nil {
			helper.LogError(apiCallID, "ID or SIM Number alerady exists: "+payload.IDNumber+" or "+payload.SIMNumber)
			return errors.New("id and sim number already exists")
		}

		createCustomer := model.Customer{
			Name:      payload.Name,
			IDNumber:  payload.IDNumber,
			SIMNumber: payload.SIMNumber,
			Phone:     payload.Phone,
			Address:   payload.Address,
			Status:    constant.CustomerStatusActive,
		}

		_, err = c.CustomerRepository.CreateCustomer(tx, createCustomer)
		if err != nil {
			helper.LogError(apiCallID, "Error creating Customer : "+err.Error())
			return errors.New("error creating customer")
		}

		return nil
	})
	if err != nil {
		switch err.Error() {
		case "id and sim number already exists":
			return constant.Res400IDOrSIMNumberExists
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Save
}

func (c *CustomerService) UpdateCustomerDetail(apiCallID string, customerUUID string, payload request.UpdateCustomerDetailRequest) constant.ResponseMap {
	err := c.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		customer, err := c.CustomerRepository.GetCustomerByUUID(tx, customerUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Customer not found: "+customerUUID)
				return errors.New("customer not found")
			}
			helper.LogError(apiCallID, "Error getting customer by UUID: "+err.Error())
			return errors.New("error getting customer by uuid")
		}

		if customer.IDNumber != payload.IDNumber || customer.SIMNumber != payload.SIMNumber {
			checkIDNumberSIMNumber, err := c.CustomerRepository.GetCustomerByIDNumberORSIMNumber(tx, payload.IDNumber, payload.SIMNumber, customerUUID)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Error checking customer ID and SIM number: "+err.Error())
				return errors.New("error get motorcycle by customer ID and SIM number")
			}

			if checkIDNumberSIMNumber != nil {
				helper.LogError(apiCallID, "ID or SIM number already exists: "+payload.IDNumber+" or "+payload.SIMNumber)
				return errors.New("id or sim number already exists")
			}
		}

		blacklistReason := ""
		if payload.Status == string(constant.CustomerStatusBlacklisted) {
			blacklistReason = payload.BlacklistReason
		}

		updateCustomer := map[string]interface{}{
			"name":             payload.Name,
			"id_number":        payload.IDNumber,
			"sim_number":       payload.SIMNumber,
			"phone":            payload.Phone,
			"address":          payload.Address,
			"status":           constant.CustomerStatus(payload.Status),
			"blacklist_reason": blacklistReason,
		}

		err = c.CustomerRepository.UpdateCustomerMap(tx, *customer, updateCustomer)
		return nil

	})

	if err != nil {
		switch err.Error() {
		case "customer not found":
			return constant.Res400CustomerNotFound
		case "id or sim number already exists":
			return constant.Res400IDOrSIMNumberExists
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Update
}

func (c *CustomerService) UpdateCustomerStatus(apiCallID string, customerUUID string, payload request.UpdateCustomerStatusRequest) constant.ResponseMap {
	err := c.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		customer, err := c.CustomerRepository.GetCustomerByUUID(tx, customerUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Customer not found: "+customerUUID)
				return errors.New("customer not found")
			}
			helper.LogError(apiCallID, "Error getting customer by UUID: "+err.Error())
			return errors.New("error getting customer by uuid")
		}

		blacklistReason := ""
		if payload.Status == string(constant.CustomerStatusBlacklisted) {
			blacklistReason = payload.BlacklistReason
		}

		updateCustomer := map[string]interface{}{
			"status":           constant.CustomerStatus(payload.Status),
			"blacklist_reason": blacklistReason,
		}

		err = c.CustomerRepository.UpdateCustomerMap(tx, *customer, updateCustomer)
		return nil
	})

	if err != nil {
		switch err.Error() {
		case "customer not found":
			return constant.Res400CustomerNotFound
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Update
}

func (c *CustomerService) GetListCustomerPagination(apiCallID string, param helper.PaginationParam, filter request.GetCustomerListFilter) ([]response.CustomerListPaginationResponse, *helper.ResponseMeta, constant.ResponseMap) {
	customers, meta, err := c.CustomerRepository.GetListCustomerPagination(c.DB, param, filter)
	if err != nil {
		helper.LogError(apiCallID, "Error getting customer list: "+err.Error())
		return nil, nil, constant.Res422SomethingWentWrong
	}

	formatedCustomerList := response.CustomerListPaginationFormatter(customers)
	return formatedCustomerList, &meta, constant.Res200Get
}

func (c *CustomerService) GetDetailCustomer(apiCallID string, customerUUID string) (*response.CustomerDetailResponse, constant.ResponseMap) {
	customer, err := c.CustomerRepository.GetCustomerByUUID(c.DB, customerUUID, false)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogError(apiCallID, "Customer not found")
			return nil, constant.Res400CustomerNotFound
		}
		helper.LogError(apiCallID, "Unable to get customer with error : "+err.Error())
		return nil, constant.Res422SomethingWentWrong
	}
	formattedCustomerDetail := response.CustomerDetailFormatter(*customer)
	return &formattedCustomerDetail, constant.Res200Get
}
