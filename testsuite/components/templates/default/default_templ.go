// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package defaulttemplate

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb"
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar/navitem"
)

type DefaultTemplateProps struct {
	SiteNavigation            templ.Component
	BreadCrumbItems           []*breadcrumbitem.BreadcrumbItemProps
	ComponentsNavItems        []*navitem.NavItemProps
	ComponentsNavItemsSidebar []navitem.NavItemProps
}

var components = []string{
	"accordion",
	"alerts",
	"badge",
	"breadcrumb",
	"card",
}

func (dtp DefaultTemplateProps) GetComponentLinksForMainNavigation() (items []*navitem.NavItemProps) {

	for _, componentName := range components {
		dtp.ComponentsNavItems = append(dtp.ComponentsNavItems, &navitem.NavItemProps{
			Name: componentName,
			Link: fmt.Sprintf("/components/%s", componentName),
		})
	}
	items = dtp.ComponentsNavItems
	return

}

func Show(props DefaultTemplateProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container-fluid\"><div class=\"row mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = navbar.Show(navbar.NavbarProps{
			Brand:    "TemplStrap.io",
			NavItems: props.ComponentsNavItems,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"row\"><!--")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := ` Left Column `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("--><div class=\"col-md-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.BreadCrumbItems != nil {
			templ_7745c5c3_Err = breadcrumb.Show(breadcrumb.BreadcrumbProps{
				BreadcrumbItems: props.BreadCrumbItems,
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}