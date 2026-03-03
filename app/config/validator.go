package config

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
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

	err = validate.RegisterValidation("date_must_before", DateMustBefore)
	if err != nil {
		log.Println("[ERROR] Error register validation date_must_before : " + err.Error())
	}

	err = validate.RegisterValidation("date_must_after", DateMustAfter)
	if err != nil {
		log.Println("[ERROR] Error register validation date_must_after : " + err.Error())
	}

	err = validate.RegisterValidation("string_number", StringOfNumber)
	if err != nil {
		log.Println("[ERROR] Error register validation string_number : " + err.Error())
	}

	err = validate.RegisterValidation("time_must_before", TimeMustBefore)
	if err != nil {
		log.Println("[ERROR] Error register validation time_must_before : " + err.Error())
	}

	err = validate.RegisterValidation("time_must_after", TimeMustAfter)
	if err != nil {
		log.Println("[ERROR] Error register validation time_must_after : " + err.Error())
	}

	err = validate.RegisterValidation("unique_combination", UniqueCombination)
	if err != nil {
		log.Println("[ERROR] Error register validation unique_combination : " + err.Error())
	}

	err = validate.RegisterValidation("date_only_must_before", DateOnlyMustBefore)
	if err != nil {
		log.Println("[ERROR] Error register validation date_only_must_before : " + err.Error())
	}

	err = validate.RegisterValidation("date_only_must_after", DateOnlyMustAfter)
	if err != nil {
		log.Println("[ERROR] Error register validation date_only_must_after : " + err.Error())
	}

	err = validate.RegisterValidation("time_must_before", TimeMustBefore)
	if err != nil {
		log.Println("[ERROR] Error register validation time_must_before : " + err.Error())
	}

	err = validate.RegisterValidation("time_must_after", TimeMustAfter)
	if err != nil {
		log.Println("[ERROR] Error register validation time_must_after : " + err.Error())
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

	err = validate.RegisterValidation("future_date_or_today", FutureDateOrToday)
	if err != nil {
		log.Println("[ERROR] Error register validation future_date_or_today : " + err.Error())
	}

	err = validate.RegisterValidation("url_or_path", URLOrPath)
	if err != nil {
		log.Println("[ERROR] Error register validation url_or_path : " + err.Error())
	}
	err = validate.RegisterValidation("cms_admin_password", CMSAdminPasswordValidation)
	if err != nil {
		log.Println("[ERROR] Error register validation cms_admin_password : " + err.Error())
	}

	err = validate.RegisterValidation("plate_number", AlphanumericOnly)
	if err != nil {
		log.Println("[ERROR] Error register validation plate_number : " + err.Error())
	}

	return validate
}

func NotOnlySpace(fl validator.FieldLevel) bool {
	data := fl.Field().String()

	return strings.TrimSpace(data) != ""
}

func CMSAdminPasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Parent().FieldByName("Password").String()

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[^a-zA-Z\d]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	return hasUppercase && hasLowercase && hasSymbol && hasDigit
}

func DateMustBefore(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() {
		return false
	}

	firstDate, err := time.Parse(time.DateTime, fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := time.Parse(time.DateTime, secondDateField.String())
	if err != nil {
		return false
	}

	return firstDate.Before(secondDate)
}

func DateMustAfter(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() {
		return false
	}

	firstDate, err := time.Parse(time.DateTime, fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := time.Parse(time.DateTime, secondDateField.String())
	if err != nil {
		return false
	}

	return firstDate.After(secondDate)
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

func parseFlexibleTime(value string) (time.Time, error) {
	if t, err := time.Parse(time.TimeOnly, value); err == nil {
		return t, nil
	}

	return time.Parse("15:04", value)
}

func TimeMustBefore(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() {
		return false
	}

	firstDate, err := parseFlexibleTime(fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := parseFlexibleTime(secondDateField.String())
	if err != nil {
		return false
	}

	return firstDate.Before(secondDate)
}

func TimeMustAfter(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() {
		return false
	}

	firstDate, err := parseFlexibleTime(fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := parseFlexibleTime(secondDateField.String())
	if err != nil {
		return false
	}

	return firstDate.After(secondDate)
}

func UniqueCombination(fl validator.FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return false
	}

	field := fl.Field()
	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}

	seen := make(map[string]struct{})

	for i := 0; i < field.Len(); i++ {
		item := field.Index(i)

		if item.Kind() != reflect.Struct {
			continue
		}

		// get the target field dynamically
		v := item.FieldByName(param)
		if !v.IsValid() {
			// fallback: case-insensitive match
			for j := 0; j < item.NumField(); j++ {
				if strings.EqualFold(item.Type().Field(j).Name, param) {
					v = item.Field(j)
					break
				}
			}
		}
		if !v.IsValid() {
			return false
		}

		// flatten field to string (supports slice or scalar)
		var key string
		if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			var parts []string
			for j := 0; j < v.Len(); j++ {
				parts = append(parts, fmt.Sprintf("%v", v.Index(j).Interface()))
			}
			sort.Strings(parts)
			key = strings.Join(parts, "|")
		} else {
			key = fmt.Sprintf("%v", v.Interface())
		}

		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}

	return true
}

func ProductNameValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9 .&-]+$`)
	return re.MatchString(value)
}

func SKUValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(value)
}

func DateOnlyMustBefore(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() || secondDateField.String() == "" {
		return true
	}

	firstDate, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := time.Parse("2006-01-02", secondDateField.String())
	if err != nil {
		return true
	}

	return firstDate.Before(secondDate)
}

func DateOnlyMustAfter(fl validator.FieldLevel) bool {
	secondDateField := fl.Parent().FieldByName(fl.Param())
	if !secondDateField.IsValid() || secondDateField.String() == "" {
		return true
	}

	firstDate, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	secondDate, err := time.Parse("2006-01-02", secondDateField.String())
	if err != nil {
		return true
	}

	return firstDate.After(secondDate)
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

func FutureDateOrToday(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	if dateStr == "" {
		return false
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.Local
	}

	parsedDate, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return false
	}

	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	return !parsedDate.Before(today)
}

func URLOrPath(fl validator.FieldLevel) bool {
	value := strings.TrimSpace(fl.Field().String())
	if value == "" {
		return true
	}
	urlRegex := regexp.MustCompile(`^(http|https)://[^\s/$.?#].[^\s]*$`)
	if urlRegex.MatchString(value) {
		return true
	}
	pathRegex := regexp.MustCompile(`^[A-Za-z0-9/_\-.]+$`)
	return pathRegex.MatchString(value)
}
