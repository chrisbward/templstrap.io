package validationcontroller

import (
	"github.com/chrisbward/templstrap.io/testsuite/entities"
)

//go:generate go run go.uber.org/mock/mockgen -package mock_validationcontroller -destination=./mocks/mock_validationcontroller.go -source=./validationcontroller.go -typed
type IValidationController interface {
	DoValidation() (validationPassed bool, fieldValidators []entities.FormField)
	InvalidateFields()
}

type ValidationController struct {
	FieldValidators map[string]*entities.FormField
}

func NewValidationController(fieldValidators map[string]*entities.FormField) IValidationController {
	return &ValidationController{
		FieldValidators: fieldValidators,
	}
}

func (vc *ValidationController) DoValidation() (validationPassed bool, fieldValidators []entities.FormField) {

	validationPassed = true

	for _, formField := range vc.FieldValidators {
		formField.FieldValidationState = entities.FieldIsValidValidationStatusType
		isValid := formField.ValidateField()
		if !isValid {
			formField.FieldValidationState = entities.FieldIsInvalidValidationStatusType
			validationPassed = false
		}
	}

	return
}

func (vc *ValidationController) InvalidateFields() {
	for _, formField := range vc.FieldValidators {
		formField.FieldValidationState = entities.FieldIsInvalidValidationStatusType
		formField.FieldValue = ""
	}
}
