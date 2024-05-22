// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package navsandtabs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/chrisbward/templstrap.io/pkg/base"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar/navitem"
	"strings"
)

const RootClassName = "nav"

type NavTabStyle string

const TabStyle NavTabStyle = "nav-tabs"
const PillStyle NavTabStyle = "nav-pills"
const UnderlineStyle NavTabStyle = "nav-underline"

type HorizontalAlignment string

const AlignDefault HorizontalAlignment = ""
const AlignCenter HorizontalAlignment = "justify-content-center"
const AlignEnd HorizontalAlignment = "justify-content-end"

type NavsAndTabsProps struct {
	base.FormElementProps
	NavItems            []*navitem.NavItemProps
	TabStyle            NavTabStyle
	IsVertical          bool
	HorizontalAlignment HorizontalAlignment
	ExtraAttributes     map[string]any
}

func (natp NavsAndTabsProps) BuildClassNames() (classes string) {
	classNames := []string{RootClassName}

	classNames = append(classNames, string(natp.TabStyle))
	classNames = append(classNames, string(natp.HorizontalAlignment))

	if natp.IsVertical {
		classNames = append(classNames, "flex-column")
	}

	classes = strings.Join(classNames, " ")

	return
}

func Show(props NavsAndTabsProps) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{props.BuildClassNames()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/navsandtabs/navsandtabs.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, NavItem := range props.NavItems {
			templ_7745c5c3_Err = navitem.Show(navitem.NavItemProps{
				Name:          NavItem.Name,
				Link:          NavItem.Link,
				IsActive:      NavItem.IsActive,
				IsDisabled:    NavItem.IsDisabled,
				IsDropdown:    NavItem.IsDropdown,
				ChildNavItems: NavItem.ChildNavItems,
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
