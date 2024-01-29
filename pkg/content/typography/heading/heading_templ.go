// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package heading

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/chrisbward/templstrap.io/pkg/base"
	"strings"
)

type HeadingClass string

const h1HeadingClass HeadingClass = "h1"
const h2HeadingClass HeadingClass = "h2"
const h3HeadingClass HeadingClass = "h3"
const h4HeadingClass HeadingClass = "h4"
const h5HeadingClass HeadingClass = "h5"
const h6HeadingClass HeadingClass = "h6"

type DisplayClass string

const display1Class DisplayClass = "display-1"
const display2Class DisplayClass = "display-2"
const display3Class DisplayClass = "display-3"
const display4Class DisplayClass = "display-4"
const display5Class DisplayClass = "display-5"
const display6Class DisplayClass = "display-6"

type HeadingProps struct {
	base.ElementProps
	HeadingClass HeadingClass
	DisplayClass DisplayClass
	IsLead       bool
}

func (hp *HeadingProps) CalculateClassNames() (classNames string) {

	classValues := []string{
		string(hp.HeadingClass),
		string(hp.DisplayClass),
	}

	if hp.IsLead {
		classValues = append(classValues, "lead")
	}

	if len(hp.AdditionalClasses) > 0 {
		classValues = append(classValues, hp.AdditionalClasses...)
	}

	classNames = strings.Join(classValues, " ")

	return
}

func Show(props HeadingProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{props.CalculateClassNames()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}