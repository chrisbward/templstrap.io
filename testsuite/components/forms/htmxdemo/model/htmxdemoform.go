package model

import (
	"github.com/a-h/templ"
	htmxdemoform "github.com/chrisbward/templstrap.io/testsuite/components/forms/htmxdemo/view"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
)

type HTMXDemoForm struct {
	FormFields           map[string]*entities.FormField
	FormFieldOrder       []string
	IncorrectCredentials bool
}

func (hdf *HTMXDemoForm) GetForm() (template templ.Component) {
	return htmxdemoform.Show(htmxdemoform.HTMXDemoFormProps{
		FormFields:           hdf.FormFields,
		FormFieldOrder:       hdf.FormFieldOrder,
		IncorrectCredentials: hdf.IncorrectCredentials,
	})
}
