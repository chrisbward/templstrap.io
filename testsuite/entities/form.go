package entities

import (
	"github.com/a-h/templ"
	"github.com/chrisbward/templstrap.io/pkg/base"
)

//go:generate go run go.uber.org/mock/mockgen -package mock_formcontroller -destination=../controllers/formcontroller/mocks/mock_formcontroller.go -source=../controllers/formcontroller/formcontroller.go -typed
type IFormController interface {
	GetFormComponent() templ.Component
	DoValidation() templ.Component
}

// type FormFields struct {
// 	Name      string
// 	InputType string
// }

type InputType string

const TextType InputType = "text"
const NumberType InputType = "number"
const PlainTextType InputType = "plaintext"
const PasswordType InputType = "password"
const ColorType InputType = "color"
const DatalistType InputType = "datalist"
const FileType InputType = "file"
const EmailType InputType = "email"
const HiddenType InputType = "hidden"
const BooleanType InputType = "boolean"
const TextareaType InputType = "textarea"

type FormFieldType struct {
	InputType InputType
	Settings  map[string]interface{}
}

type FormField struct {
	Name                           string
	FieldName                      string
	FieldLabel                     string
	FieldPlaceholder               string
	FieldType                      FormFieldType
	FieldValidationState           ValidationStateType
	FieldValidationFeedbackMessage string
	FieldValue                     string
	FieldListItems                 []base.FormListProps
	FieldValidationType            []ValidationType
}

type Form struct {
	Fields []FormField
	// ValidationController validationcontroller.IValidationController
	SucessTemplate templ.Component
}

type ValidationStateType string

const FieldIsInvalidValidationStatusType ValidationStateType = "is-invalid"
const FieldIsValidValidationStatusType ValidationStateType = "is-valid"

type ValidationType string

const FieldEmptyIsInvalid ValidationType = "empty-field-invalid"
const FieldRegexMatchIsInvalid ValidationType = "regex-field-invalid"

func (ff *FormField) ValidateField() (FieldIsValid bool) {
	FieldIsValid = true
	for _, validationType := range ff.FieldValidationType {
		// empty field
		if validationType == FieldEmptyIsInvalid && ff.FieldValue == "" {
			return false
		}
	}

	return
}

type FormFieldValidation struct {
	IsRequired      bool
	ValidationState ValidationStateType
}
