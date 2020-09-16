package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidDynamoDBTable(fl validator.FieldLevel) bool {
	table := fl.Field().String()

	if len(table) < 3 || len(table) > 255 {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9_.-]+$", table)

	return match
}
