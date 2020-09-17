package test

import (
	"strings"
	"testing"

	validator "github.com/hexbee-net/aws-validator"
	"github.com/stretchr/testify/assert"
)

type DynamoDBTest struct {
	DynamoTableNameValue string `validate:"dynamodb"`
}

func TestIsValidDynamoDBTable(t *testing.T) {
	cases := []struct {
		input string
		err   string
	}{
		{
			input: "no",
			err:   "Key: 'DynamoDBTest.DynamoTableNameValue' Error:Field validation for 'DynamoTableNameValue' failed on the 'dynamodb' tag",
		},
		{
			input: strings.Repeat("a", 256),
			err:   "Key: 'DynamoDBTest.DynamoTableNameValue' Error:Field validation for 'DynamoTableNameValue' failed on the 'dynamodb' tag",
		},
		{
			input: "foo bar",
			err:   "Key: 'DynamoDBTest.DynamoTableNameValue' Error:Field validation for 'DynamoTableNameValue' failed on the 'dynamodb' tag",
		},
		{
			input: "ABORT",
			err:   "Key: 'DynamoDBTest.DynamoTableNameValue' Error:Field validation for 'DynamoTableNameValue' failed on the 'dynamodb' tag",
		},
		{
			input: "TableName",
		},
		{
			input: "TableName01",
		},
		{
			input: "Table.Name",
		},
		{
			input: "Table_Name",
		},
		{
			input: "Table-Name",
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			v := validator.New()
			err := v.Struct(DynamoDBTest{DynamoTableNameValue: tc.input})

			if tc.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.err)
			}
		})
	}
}
