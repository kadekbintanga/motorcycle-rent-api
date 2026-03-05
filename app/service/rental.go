package service

import (
	"errors"
	"fmt"
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RentalServiceInterface interface {
	CreateRental(apiCallID string, payload request.CreateRentalRequest) (*response.CreateRentalResponse, constant.ResponseMap)
	ReturnRental(apiCallID string, rentalUUID string, payload request.ReturnRentalRequest) constant.ResponseMap
}

type RentalService struct {
	DB                   *gorm.DB
	RentalRepository     repository.RentalRepositoryInterface
	CustomerRepository   repository.CustomerRepositoryInterface
	MotorcycleRepository repository.MotorcycleRepositoryInterface
	PaymentRepository    repository.PaymentRepositoryInterface
}

func NewRentalService(db *gorm.DB, rentalRepository repository.RentalRepositoryInterface, customerRepository repository.CustomerRepositoryInterface, motorcycleRepository repository.MotorcycleRepositoryInterface, paymentRepository repository.PaymentRepositoryInterface) RentalServiceInterface {
	return &RentalService{
		DB:                   db,
		RentalRepository:     rentalRepository,
		CustomerRepository:   customerRepository,
		MotorcycleRepository: motorcycleRepository,
		PaymentRepository:    paymentRepository,
	}
}

func (r *RentalService) CreateRental(apiCallID string, payload request.CreateRentalRequest) (*response.CreateRentalResponse, constant.ResponseMap) {
	var formattedCreateRental response.CreateRentalResponse
	err := r.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		customer, err := r.CustomerRepository.GetCustomerByUUID(tx, payload.CustomerUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Customer not found: "+payload.CustomerUUID)
				return errors.New("customer not found")
			}
			helper.LogError(apiCallID, "Error getting customer by UUID: "+err.Error())
			return errors.New("error getting customer by uuid")
		}

		if customer.Status == constant.CustomerStatusBlacklisted {
			helper.LogError(apiCallID, "Customer was blacklisted")
			return errors.New("customer blacklisted")
		}

		motorcycle, err := r.MotorcycleRepository.GetMotorcycleByUUID(tx, payload.MotorcycleUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Motorcycle not found: "+payload.MotorcycleUUID)
				return errors.New("motocycle not found")
			}
			helper.LogError(apiCallID, "Error getting motorcycle by UUID: "+err.Error())
			return errors.New("error getting motorcycle by uuid")
		}

		if motorcycle.Status != constant.MotorcycleStatusAvailable {
			helper.LogError(apiCallID, "motorcycle is not available")
			return errors.New("motor unavailable")
		}

		checkCustomerRental, err := r.RentalRepository.GetCustomerOngoingRental(tx, payload.CustomerUUID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogError(apiCallID, "Error check customer ongoing rental: "+err.Error())
			return errors.New("error get check customer ongoing rental")
		}

		if checkCustomerRental != nil {
			helper.LogError(apiCallID, "Customer has been had ongiong rental")
			return errors.New("customer ongoing rental")
		}

		rentPrice := motorcycle.PricePerDay * float64(payload.Days)

		if payload.PaymentType == string(constant.PaymentTypeRentPayment) {
			if payload.Payment != rentPrice {
				return errors.New("invalid rent price payment")
			}
		} else {
			if payload.Payment >= rentPrice {
				return errors.New("deposit greater than rent price")
			}
		}

		rentDate := time.Now()
		returnDatePlan := rentDate.AddDate(0, 0, int(payload.Days))

		createRental := model.Rental{
			CustomerUUID:                  customer.UUID,
			MotorcycleUUID:                motorcycle.UUID,
			RentDate:                      rentDate,
			ReturnDatePlan:                returnDatePlan,
			PricePerDayCaptured:           motorcycle.PricePerDay,
			CustomerNameCaptured:          customer.Name,
			CustomerIDNumberCaptured:      customer.IDNumber,
			CustomerSIMNumberCaptured:     customer.SIMNumber,
			CustomerPhoneCaptured:         customer.Phone,
			MotorcyclePlateNumberCaptured: motorcycle.PlateNumber,
			Status:                        constant.RentalStatusOngoing,
			RentPrice:                     rentPrice,
		}

		saveRental, err := r.RentalRepository.CreateRental(tx, createRental)
		if err != nil {
			helper.LogError(apiCallID, "Error creating Rental : "+err.Error())
			return errors.New("error creating rental")
		}

		payment := model.Payment{
			RentalUUID: saveRental.UUID,
			Amount:     payload.Payment,
			Type:       constant.PaymentType(payload.PaymentType),
			Method:     constant.PaymentMethod(payload.PaymentMethod),
		}

		_, err = r.PaymentRepository.CreatePayment(tx, payment)
		if err != nil {
			helper.LogError(apiCallID, "Error creating Payment : "+err.Error())
			return errors.New("error creating payment")
		}

		updateMotorcycle := map[string]interface{}{
			"status": constant.MotorcycleStatusRented,
		}

		err = r.MotorcycleRepository.UpdateMotorcycleMap(tx, *motorcycle, updateMotorcycle)
		if err != nil {
			helper.LogError(apiCallID, "Error update motorcycle status : "+err.Error())
			return errors.New("error update motorcycle status")
		}

		formattedCreateRental = response.CreateRentalResponseFormatter(*saveRental)
		return nil
	})
	if err != nil {
		switch err.Error() {
		case "customer not found":
			return nil, constant.Res400CustomerNotFound
		case "customer blacklisted":
			return nil, constant.Res400CustomerBlacklisted
		case "motocycle not found":
			return nil, constant.Res400MotorcycleNotFound
		case "motor unavailable":
			return nil, constant.Res400MotorcycleUnavailable
		case "customer ongoing rental":
			return nil, constant.Res400CustomerOngoingRent
		case "invalid rent price payment":
			return nil, constant.Res400InvalidRentPrice
		case "deposit greater than rent price":
			return nil, constant.Res400DepositGreaterThanRentPrice
		default:
			return nil, constant.Res422SomethingWentWrong
		}
	}
	return &formattedCreateRental, constant.Res200Save
}

func (r *RentalService) ReturnRental(apiCallID string, rentalUUID string, payload request.ReturnRentalRequest) constant.ResponseMap {
	var remainingPayment float64
	err := r.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		rental, err := r.RentalRepository.GetRentalByUUID(tx, rentalUUID, true)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Rental not found: "+rentalUUID)
				return errors.New("rental not found")
			}
			helper.LogError(apiCallID, "Error getting rental by UUID: "+err.Error())
			return errors.New("error getting rental by uuid")
		}

		if rental.Status != constant.RentalStatusOngoing {
			helper.LogError(apiCallID, "Rental has been returned")
			return errors.New("rental done")
		}

		returnDateActual, err := time.Parse("2006-01-02", payload.ReturnDate)
		if err != nil {
			return errors.New("invalid return date")
		}
		timeDiff := returnDateActual.Sub(rental.ReturnDatePlan)
		penaltyDays := int(timeDiff.Hours() / 24)

		actualPrice := rental.RentPrice
		var penaltyPrice float64

		if penaltyDays > 0 {
			penaltyPrice = float64(penaltyDays) * rental.PricePerDayCaptured
			actualPrice = rental.RentPrice + penaltyPrice
		}

		var lastPayment float64

		for _, payment := range rental.Payment {
			lastPayment = lastPayment + payment.Amount
		}

		helper.LogInfo(apiCallID, "penaltyDay : ")
		helper.LogInfo(apiCallID, penaltyDays)

		helper.LogInfo(apiCallID, "penaltyPrice : ")
		helper.LogInfo(apiCallID, penaltyPrice)

		helper.LogInfo(apiCallID, "lastPayment : ")
		helper.LogInfo(apiCallID, lastPayment)

		remainingPayment = actualPrice - lastPayment
		helper.LogInfo(apiCallID, "remaingPrice : ")
		helper.LogInfo(apiCallID, remainingPayment)

		if remainingPayment != payload.Payment {
			helper.LogError(apiCallID, "Payment amount is not equal remaining payment")
			return errors.New("invalid payment amount")
		}

		updateRental := map[string]interface{}{
			"return_date_actual": returnDateActual,
			"late_day":           penaltyDays,
			"penalty_price":      penaltyPrice,
			"status":             constant.RentalStatusCompleted,
		}

		err = r.RentalRepository.UpdateRentalMap(tx, *rental, updateRental)
		if err != nil {
			helper.LogError(apiCallID, "Error update rental data : "+err.Error())
			return errors.New("error update rental data")
		}

		if payload.Payment > 0 {
			payment := model.Payment{
				RentalUUID: rental.UUID,
				Amount:     payload.Payment,
				Type:       constant.PaymentTypeFullPayment,
				Method:     constant.PaymentMethod(payload.PaymentMethod),
			}

			_, err = r.PaymentRepository.CreatePayment(tx, payment)
			if err != nil {
				helper.LogError(apiCallID, "Error creating Payment : "+err.Error())
				return errors.New("error creating payment")
			}
		}

		updateMotorcycle := model.Motorcycle{
			Status: constant.MotorcycleStatusAvailable,
		}

		err = r.MotorcycleRepository.UpdateMotocycleByUUID(tx, rental.MotorcycleUUID.String(), updateMotorcycle)
		if err != nil {
			helper.LogError(apiCallID, "Error update motocycle : "+err.Error())
			return errors.New("error update motocycle")
		}

		if penaltyDays >= global.GlobalConfig.BlacklistDayLate {
			blacklistReason := fmt.Sprintf("Late return %d days", penaltyDays)
			updateCustomer := model.Customer{
				Status:          constant.CustomerStatusBlacklisted,
				BlacklistReason: blacklistReason,
			}
			err = r.CustomerRepository.UpdateCustomerByUUID(tx, rental.CustomerUUID.String(), updateCustomer)
			if err != nil {
				helper.LogError(apiCallID, "Error update customer : "+err.Error())
				return errors.New("error update customer")
			}
		}
		return nil
	})
	if err != nil {
		switch err.Error() {
		case "rental not found":
			return constant.Res400RentalNotFound
		case "invalid return date":
			return constant.Res400InvalidReturnDate
		case "invalid payment amount":
			message := fmt.Sprintf("Payment should be %.2f", remainingPayment)
			return constant.ResponseMap{Code: http.StatusBadRequest, Status: constant.ResponseStatusFailed, Message: message}
		case "rental done":
			return constant.Res400RentalHasBeenDone
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Update
}
