package repository

import (
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	CreateCustomer(db *gorm.DB, customer model.Customer) (*model.Customer, error)
	GetCustomerByIDNumberORSIMNumber(db *gorm.DB, idNumber string, simNumber string, excludeCustomerUUID string) (*model.Customer, error)
	UpdateCustomerMap(db *gorm.DB, customer model.Customer, updateData map[string]interface{}) error
	GetCustomerByUUID(db *gorm.DB, customerUUID string, withPreload bool) (*model.Customer, error)
	GetListCustomerPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetCustomerListFilter) ([]model.Customer, helper.ResponseMeta, error)
	UpdateCustomerByUUID(db *gorm.DB, customerUUID string, updateData model.Customer) error
	GetCustomerSummary(db *gorm.DB) (*response.CustomerSummary, error)
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

	if err := query.Order("name ASC").Find(&customer).Error; err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	meta := helper.ResponseMeta{
		Page:  param.Page,
		Limit: param.Limit,
		Total: int(total),
	}

	return customer, meta, nil
}

func (c *CustomerRepository) UpdateCustomerByUUID(db *gorm.DB, customerUUID string, updateData model.Customer) error {
	err := db.Model(&model.Customer{}).Where("uuid = ?", customerUUID).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) GetCustomerSummary(db *gorm.DB) (*response.CustomerSummary, error) {
	var result response.CustomerSummary
	err := db.Table("customers").
		Select(`
			COUNT(*) as total,
			COALESCE(SUM(CASE WHEN status = 'ACTIVE' THEN 1 ELSE 0 END),0) as active,
			COALESCE(SUM(CASE WHEN status = 'BLACKLISTED' THEN 1 ELSE 0 END),0) as blacklisted
		`).
		Where("deleted_at IS NULL").Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
