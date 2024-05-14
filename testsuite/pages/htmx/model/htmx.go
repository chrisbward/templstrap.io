package htmxpagemodel

import (
	"github.com/a-h/templ"
	"github.com/chrisbward/templstrap.io/testsuite/components/forms/htmxdemo/model"
	"github.com/chrisbward/templstrap.io/testsuite/components/htmx/loginthanks"
	"github.com/chrisbward/templstrap.io/testsuite/controllers/validationcontroller"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type HTMXPageModel struct {
	Ctx        echo.Context
	IsPostback bool
	FieldOrder []string
	Fields     map[string]*entities.FormField
}

func NewHTMXPageModel(ctx echo.Context, isPostback bool) HTMXPageModel {

	formFields := map[string]*entities.FormField{
		"txbUsername": {
			Name:             "txbUsername",
			FieldName:        "txbUsername",
			FieldLabel:       "Username",
			FieldPlaceholder: "Enter your username",
			FieldType: entities.FormFieldType{
				InputType: entities.TextType,
			},
			FieldValue: "",
			FieldValidationType: []entities.ValidationType{
				entities.FieldEmptyIsInvalid,
			},
		},
		"txbPassword": {
			Name:             "txbPassword",
			FieldName:        "txbPassword",
			FieldLabel:       "Password",
			FieldPlaceholder: "Enter your password",
			FieldType: entities.FormFieldType{
				InputType: entities.TextType,
			},
			FieldValue: "",
			FieldValidationType: []entities.ValidationType{
				entities.FieldEmptyIsInvalid,
			},
		},
	}

	if isPostback {

		// Parse the form data
		err := ctx.Request().ParseForm()
		if err != nil {
			logrus.Warnln(err.Error())
		}

		// Create a map to store form fields
		// fields := make(map[string]string)

		// Iterate through form values
		for key, values := range ctx.Request().Form {
			// For simplicity, assuming a single value for each field
			if len(values) > 0 {
				// fields[key] = values[0]
				if formFieldFound, ok := formFields[key]; ok {
					formFieldFound.FieldValue = values[0]
					formFields[key] = formFieldFound
				}
			}
		}
	}
	return HTMXPageModel{
		Ctx:        ctx,
		IsPostback: isPostback,
		FieldOrder: []string{"txbUsername", "txbPassword"},
		Fields:     formFields,
	}
}

func (hm *HTMXPageModel) GetForm() (formTemplate templ.Component) {

	htmxDemoForm := &model.HTMXDemoForm{
		FormFields:     hm.Fields,
		FormFieldOrder: hm.FieldOrder,
	}
	return htmxDemoForm.GetForm()
}
func (hm *HTMXPageModel) ValidateAndGetFormOrSuccess() (formTemplate templ.Component) {

	validationController := validationcontroller.NewValidationController(hm.Fields)

	validationPassed, _ := validationController.DoValidation()

	if validationPassed {
		if hm.Fields["txbUsername"].FieldValue == "admin" &&
			hm.Fields["txbPassword"].FieldValue == "admin" {
			return loginthanks.Show()
		} else {
			htmxDemoForm := &model.HTMXDemoForm{
				FormFields:           hm.Fields,
				FormFieldOrder:       hm.FieldOrder,
				IncorrectCredentials: true,
			}
			validationController.InvalidateFields()
			return htmxDemoForm.GetForm()
		}
	}

	htmxDemoForm := &model.HTMXDemoForm{
		FormFields:     hm.Fields,
		FormFieldOrder: hm.FieldOrder,
	}
	return htmxDemoForm.GetForm()
}
