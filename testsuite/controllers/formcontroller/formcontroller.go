package formcontroller

import (
	"github.com/a-h/templ"
	"github.com/chrisbward/templstrap.io/testsuite/controllers/validationcontroller"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
)

type FormController struct {
	Form entities.Form
}

func NewFormController(form entities.Form, validationController validationcontroller.IValidationController) entities.IFormController {

	// form.ValidationController = validationController

	return &FormController{
		Form: form,
	}
}

func (fc *FormController) GetFormComponent() (templComponent templ.Component) {
	return
}

func (fc *FormController) GetSuccessActionComponent() (templComponent templ.Component) {

	return
}

func (fc *FormController) DoValidation() (templComponent templ.Component) {

	// validationPassed, _ := fc.Form.ValidationController.DoValidation()

	// if validationPassed {
	// 	return fc.GetSuccessActionComponent()
	// }
	return
}
