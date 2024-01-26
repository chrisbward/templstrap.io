package validationcontroller_test

import (
	"testing"

	"github.com/chrisbward/templstrap.io/testsuite/controllers/validationcontroller"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
	"github.com/stretchr/testify/assert"
)

func TestDoValidation(t *testing.T) {
	testCases := []struct {
		Desc                                string
		Fields                              map[string]*entities.FormField
		ExpectedValidationResult            bool
		ExecptedFieldsWithInvalidStatusType []string
		ErrAssert                           func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool
		ExpectedErrorMessage                string
	}{
		{
			Desc: "TEST001 - Test validation for happy path",
			Fields: map[string]*entities.FormField{
				"txbUsername": {
					FieldName:  "txbUsername",
					FieldValue: "myusername",
					FieldValidationType: []entities.ValidationType{
						entities.FieldEmptyIsInvalid,
						entities.FieldRegexMatchIsInvalid,
					},
				},
				"txbPassword": {
					FieldName:  "txbPassword",
					FieldValue: "mypassword",
					FieldValidationType: []entities.ValidationType{
						entities.FieldEmptyIsInvalid,
					},
				},
			},
			ExpectedValidationResult:            true,
			ExecptedFieldsWithInvalidStatusType: []string{},
		},
		{
			Desc: "TEST002 - Test validation for unhappy path, one of both fields invalid",
			Fields: map[string]*entities.FormField{
				"txbUsername": {
					FieldName:  "txbUsername",
					FieldValue: "myusername",
					FieldValidationType: []entities.ValidationType{
						entities.FieldEmptyIsInvalid,
					},
				},
				"txbPassword": {
					FieldName:  "txbPassword",
					FieldValue: "",
					FieldValidationType: []entities.ValidationType{
						entities.FieldEmptyIsInvalid,
					},
				},
			},
			ExpectedValidationResult:            false,
			ExecptedFieldsWithInvalidStatusType: []string{"txbPassword"},
		},
	}
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Desc, func(t *testing.T) {
			validationController := validationcontroller.NewValidationController(tc.Fields)

			validationResult, validatedFields := validationController.DoValidation()

			for _, validatedField := range validatedFields {
				for _, expectedInvalidField := range tc.ExecptedFieldsWithInvalidStatusType {
					if expectedInvalidField == validatedField.FieldName {
						assert.Equal(t, validatedField.FieldValidationState, entities.FieldIsInvalidValidationStatusType)
					}
				}
			}

			// assert.Contains(t, validatedFields, tc.ExecptedFieldsWithInvalidStatusType, "Expected '%v' to be in the result for input %v", tc.ExecptedFieldsWithInvalidStatusType, validatedFields)

			assert.Equal(t, tc.ExpectedValidationResult, validationResult)
		})
	}
}

func TestValidateField(t *testing.T) {
	testCases := []struct {
		Desc                          string
		FormField                     entities.FormField
		ExpectedFieldValidationResult bool
	}{
		{
			Desc: "TEST001 - Test validation for happy path",
			FormField: entities.FormField{
				FieldName:  "txbUsername",
				FieldValue: "myusername",
				FieldValidationType: []entities.ValidationType{
					entities.FieldEmptyIsInvalid,
				},
			},
			ExpectedFieldValidationResult: true,
		},
		{
			Desc: "TEST002 - Test validation for happy path, no validations",
			FormField: entities.FormField{
				FieldName:  "txbUsername",
				FieldValue: "myusername",
			},
			ExpectedFieldValidationResult: true,
		},
		{
			Desc: "TEST003 - Test validation for unhappy path",
			FormField: entities.FormField{
				FieldName:  "txbUsername",
				FieldValue: "",
				FieldValidationType: []entities.ValidationType{
					entities.FieldEmptyIsInvalid,
				},
			},
			ExpectedFieldValidationResult: false,
		},
	}
	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Desc, func(t *testing.T) {

			fieldIsValid := tc.FormField.ValidateField()

			assert.Equal(t, tc.ExpectedFieldValidationResult, fieldIsValid)
		})
	}
}
