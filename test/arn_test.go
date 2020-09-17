package test

import (
	"testing"

	validator "github.com/hexbee-net/aws-validator"
	"github.com/stretchr/testify/assert"
)

type ArnTest struct {
	ArnValue string `validate:"arn"`
}

func TestIsValidARN(t *testing.T) {
	cases := []struct {
		input string
		err   string
	}{
		{
			input: "invalid",
			err:   "Key: 'ArnTest.ArnValue' Error:Field validation for 'ArnValue' failed on the 'arn' tag",
		},
		{
			input: "arn:nope",
			err:   "Key: 'ArnTest.ArnValue' Error:Field validation for 'ArnValue' failed on the 'arn' tag",
		},
		{
			input: "arn:aws:ecr:us-west-2:123456789012:repository/foo/bar",
		},
		{
			input: "arn:aws:elasticbeanstalk:us-east-1:123456789012:environment/My App/MyEnvironment",
		},
		{
			input: "arn:aws:iam::123456789012:user/David",
		},
		{
			input: "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
		},
		{
			input: "arn:aws:s3:::my_corporate_bucket/exampleobject.png",
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			v := validator.New()
			err := v.Struct(ArnTest{ArnValue: tc.input})

			if tc.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}
