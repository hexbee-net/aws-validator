package validator

import (
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/go-playground/validator/v10"
)

// IsValidObjectName verifies an object name in accordance with Amazon's
// requirements.
//
// See:
// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func IsValidARN(fl validator.FieldLevel) bool {
	return arn.IsARN(fl.Field().String())
}
