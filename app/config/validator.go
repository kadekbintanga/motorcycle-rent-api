package config

import (
	"log"
	"motorcycle-rent-api/app/constant"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func InitValidator() *validator.Validate {
	var err error
	validate := validator.New()

	// custom validation here
	err = validate.RegisterValidation("not_only_space", NotOnlySpace)
	if err != nil {
		log.Println("[ERROR] Error register validation not_only_space : " + err.Error())
	}
	err = validate.RegisterValidation("string_number", StringOfNumber)
	if err != nil {
		log.Println("[ERROR] Error register validation string_number : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_space_dot_dash_amp", AlphanumericSpaceDotDashAmp)
	if err != nil {
		log.Println("[ERROR] Error register validation alphanumeric_space_dot_dash_amp : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_space_with_symbols", AlphanumericSpaceWithSymbols)
	if err != nil {
		log.Println("[ERROR] Error register validation alphanumeric_space_with_symbols : " + err.Error())
	}

	err = validate.RegisterValidation("phone_validation", PhoneValidation)
	if err != nil {
		log.Println("[ERROR] Error register validation phone_validation : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_space_dash", AlphanumericSpaceDash)
	if err != nil {
		log.Println("[ERROR] Error register validation alphanumeric_space_dash : " + err.Error())
	}

	err = validate.RegisterValidation("numbers_only", NumbersOnly)
	if err != nil {
		log.Println("[ERROR] Error register translation numbers_only : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_dot_at_dash_underscore", AlphanumericDotAtDashUnderscore)
	if err != nil {
		log.Println("[ERROR] Error register validation alphanumeric_dot_at_dash_underscore : " + err.Error())
	}

	err = validate.RegisterValidation("alphabet_only", AlphabetOnly)
	if err != nil {
		log.Println("[ERROR] Error register translation alphabet_only : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_only", AlphanumericOnly)
	if err != nil {
		log.Println("[ERROR] Error register translation alphanumeric_only : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_dot_slash_dash", AlphanumericDotSlashDash)
	if err != nil {
		log.Println("[ERROR] Error register translation alphanumeric_dot_slash_dash : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_with_symbols", AlphanumericWithSymbols)
	if err != nil {
		log.Println("[ERROR] Error register translation alphanumeric_with_symbols : " + err.Error())
	}

	err = validate.RegisterValidation("alphanumeric_space_dash", AlphanumericSpaceDash)
	if err != nil {
		log.Println("[ERROR] Error register validation alphanumeric_space_dash : " + err.Error())
	}

	err = validate.RegisterValidation("name_validation", ValidateName)
	if err != nil {
		log.Println("[ERROR] Error register validation name : " + err.Error())
	}

	err = validate.RegisterValidation("cms_admin_password", CMSAdminPasswordValidation)
	if err != nil {
		log.Println("[ERROR] Error register validation cms_admin_password : " + err.Error())
	}

	err = validate.RegisterValidation("plate_number", AlphanumericOnly)
	if err != nil {
		log.Println("[ERROR] Error register validation plate_number : " + err.Error())
	}

	err = validate.RegisterValidation("payment_method_required", PaymentMethodRequired)
	if err != nil {
		log.Println("[ERROR] Error register validation payment_method : " + err.Error())
	}

	return validate
}

func NotOnlySpace(fl validator.FieldLevel) bool {
	data := fl.Field().String()

	return strings.TrimSpace(data) != ""
}

func PaymentMethodRequired(fl validator.FieldLevel) bool {
	payment := fl.Parent().FieldByName("Payment").Float()
	method := fl.Field().String()

	if payment > 0 {
		return method == string(constant.PaymentMethodCash) || method == string(constant.PaymentMethodQris) || method == string(constant.PaymentMethodTransfer)
	}

	return true
}

func CMSAdminPasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Parent().FieldByName("Password").String()

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[^a-zA-Z\d]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	return hasUppercase && hasLowercase && hasSymbol && hasDigit
}

func StringOfNumber(fl validator.FieldLevel) bool {
	str := fl.Field().String()

	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func AlphanumericSpaceDotDashAmp(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9 .\-&]+$`)
	return re.MatchString(value)
}

func AlphanumericSpaceWithSymbols(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9\s\p{P}\p{S}]+$`)
	return re.MatchString(value)
}

func PhoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if strings.HasPrefix(phone, "0") {
		phone = "62" + phone[1:]
	}
	re := regexp.MustCompile(`^[1-9][0-9]{7,13}$`)
	return re.MatchString(phone)
}

func AlphanumericSpaceDash(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9\- ]+$`)
	return re.MatchString(value)
}

func NumbersOnly(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return regexp.MustCompile(`^[0-9]+$`).MatchString(field.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() >= 0

	default:
		return false
	}
}

func AlphanumericDotAtDashUnderscore(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9.@\-_]+$`)
	return re.MatchString(value)
}

func AlphabetOnly(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z]+$`)
	return re.MatchString(value)
}

func AlphanumericOnly(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	return re.MatchString(value)
}

func AlphanumericDotSlashDash(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9.\-/]+$`)
	return re.MatchString(value)
}

func AlphanumericWithSymbols(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[A-Za-z0-9\p{P}\p{S}]+$`)
	return re.MatchString(value)
}

func ValidateName(fl validator.FieldLevel) bool {
	value := strings.TrimSpace(fl.Field().String())
	re := regexp.MustCompile(`^[A-Za-z '.-]+$`)
	return re.MatchString(value)
}
