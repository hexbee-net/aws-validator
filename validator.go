package validator

import "github.com/go-playground/validator/v10"

// New returns a new instance of 'validate' with sane defaults and AWS addons.
func New() *validator.Validate {
	v := validator.New()

	must(v.RegisterValidation("arn", IsValidARN))
	must(v.RegisterValidation("arn", IsValidARN))
	must(v.RegisterValidation("s3bucket", IsValidBucketName))
	must(v.RegisterValidation("s3object", IsValidObjectName))
	must(v.RegisterValidation("dynamodb", IsValidDynamoDBTable))

	return v
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
