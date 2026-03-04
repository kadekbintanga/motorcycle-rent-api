package repository

import (
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/resource/request"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	CreateCustomer(db *gorm.DB, customer model.Customer) (*model.Customer, error)
	GetCustomerByIDNumberORSIMNumber(db *gorm.DB, idNumber string, simNumber string, excludeCustomerUUID string) (*model.Customer, error)
	UpdateCustomerMap(db *gorm.DB, customer model.Customer, updateData map[string]interface{}) error
	GetCustomerByUUID(db *gorm.DB, customerUUID string, withPreload bool) (*model.Customer, error)
	GetListCustomerPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetCustomerListFilter) ([]model.Customer, helper.ResponseMeta, error)
}

type CustomerRepository struct{}

func NewCustomerRepository() CustomerRepositoryInterface {
	return &CustomerRepository{}
}

func (c *CustomerRepository) CreateCustomer(db *gorm.DB, customer model.Customer) (*model.Customer, error) {
	if err := db.Create(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) GetCustomerByIDNumberORSIMNumber(db *gorm.DB, idNumber string, simNumber string, excludeCustomerUUID string) (*model.Customer, error) {
	var customer model.Customer
	query := db.Where("id_number = ? OR sim_number = ?", idNumber, simNumber)
	if excludeCustomerUUID != "" {
		query = query.Where("uuid != ?", excludeCustomerUUID)
	}
	if err := query.First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepository) UpdateCustomerMap(db *gorm.DB, customer model.Customer, updateData map[string]interface{}) error {
	err := db.Model(customer).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) GetCustomerByUUID(db *gorm.DB, customerUUID string, withPreload bool) (*model.Customer, error) {
	var customer model.Customer
	query := db.Where("uuid = ?", customerUUID)

	if withPreload {
		query = query
	}

	err := query.First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *CustomerRepository) GetListCustomerPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetCustomerListFilter) ([]model.Customer, helper.ResponseMeta, error) {
	var customer []model.Customer

	query := db.Model(&model.Customer{})

	if param.Search != "" {
		query = query.Where("name ILIKE ? OR id_number ILIKE ? OR sim_number ILIKE ? OR phone ILIKE ?", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%")
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	query, total, err := helper.Paginate(param.Limit, param.Page, query)
	if err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	if err := query.Find(&customer).Order("name asc").Error; err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	meta := helper.ResponseMeta{
		Page:  param.Page,
		Limit: param.Limit,
		Total: int(total),
	}

	return customer, meta, nil
}
