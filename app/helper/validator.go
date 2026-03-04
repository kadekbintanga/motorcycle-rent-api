package helper

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorData struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func toSnakeCaseBuildErrorKey(s string) string {
	snake := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(s, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ValidationBuildErrorFormattedKey(namespace string) string {
	parts := strings.Split(namespace, ".")
	if len(parts) <= 1 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(namespace))
	fieldParts := parts[1:]
	for i, part := range fieldParts {
		if i > 0 {
			result.WriteRune('.')
		}
		if matches := regexp.MustCompile(`^(\w+)\[(\d+)\]$`).FindStringSubmatch(part); len(matches) == 3 {
			fieldName := matches[1]
			index := matches[2]
			result.WriteString(toSnakeCaseBuildErrorKey(fieldName))
			result.WriteRune('.')
			result.WriteString(index)
		} else {
			result.WriteString(toSnakeCaseBuildErrorKey(part))
		}
	}
	return result.String()
}

func ErrorValidationFormatter(errors validator.ValidationErrors) []ValidationErrorData {
	formattedErrors := make([]ValidationErrorData, len(errors))
	for i, err := range errors {
		errorKey := ValidationBuildErrorFormattedKey(err.StructNamespace())
		formattedErrors[i] = ValidationErrorData{
			Key:     errorKey,
			Message: ErrorValidationMessageGenerator(errorKey, err.Tag(), err.Param()),
		}
	}
	return formattedErrors
}

func ErrorValidationMessageGenerator(field, tag, param string) string {
	message := "Field validation for '" + field + "' failed on the '" + tag + "' tag"
	switch tag {
	case "required":
		message = field + " cannot be empty."
	case "oneof":
		message = "Invalid value of " + field + " one of " + param
	case "datetime":
		message = "Invalid date time format for " + field + "."
	case "required_if":
		message = field + " cannot be empty if " + param + "."
	case "min":
		message = field + " minimum value is " + param + "."
	case "max":
		message = field + " maximum value is " + param + "."
	case "len":
		message = field + " must be exactly " + param + " digits."
	case "uuid":
		message = "Invalid format of " + field + "."
	case "not_only_space":
		message = field + " value cannot be only 'space'."
	case "date_must_before":
		message = field + " must be before " + param + "."
	case "date_must_after":
		message = field + " must be after " + param + "."
	case "string_number":
		message = field + " only can be number"
	case "email":
		message = "invalid email format " + field
	case "plate_number":
		message = "Invalid plate number format."
	case "timezone":
		message = "Invalid timezone format."
	case "gte":
		message = field + " must be greater than or equal to " + param + "."
	case "lte":
		message = field + " must be less than or equal to " + param + "."
	case "latitude":
		message = "Invalid latitude format."
	case "longitude":
		message = "Invalid latitude format."
	case "unique_combination":
		message = "Each " + param + " combination must be unique"
	case "product_name":
		message = field + " only valid Alphanumeric, Space, Dot, Dash, and Ampersand"
	case "product_sku":
		message = field + " only valid Alphanumeric, Dash, and Underscore"
	case "date_only_must_before":
		message = field + " must be before " + param + "."
	case "date_only_must_after":
		message = field + " must be after " + param + "."
	case "required_with":
		message = field + " cannot be empty, required with " + param + "."
	case "cms_admin_password":
		message = field + " requires at least one uppercase, one lowercase, one symbol, and one numeric."
	case "cms_admin_new_password":
		message = field + " requires at least one uppercase, one lowercase, one symbol, and one numeric."
	case "pin":
		message = field + " must be 4 numeric digits."

	}

	return message
}

type LocalizedText struct {
	ID string `json:"id" validate:"required"`
	EN string `json:"en" validate:"required"`
}
