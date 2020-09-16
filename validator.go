package validator

import "github.com/go-playground/validator/v10"

// New returns a new instance of 'validate' with sane defaults and AWS addons.
func New() *validator.Validate {
	validate := validator.New()

	must(validate.RegisterValidation("arn", IsValidARN))
	must(validate.RegisterValidation("arn", IsValidARN))
	must(validate.RegisterValidation("s3bucket", IsValidBucketName))
	must(validate.RegisterValidation("s3object", IsValidObjectName))
	must(validate.RegisterValidation("dynamodb", IsValidDynamoDBTable))

	return validate
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
