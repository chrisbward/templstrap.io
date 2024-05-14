// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package breadcrumb

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/chrisbward/templstrap.io/pkg/base"
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
)

type BreadcrumbProps struct {
	base.ElementProps
	BreadcrumbItems []*breadcrumbitem.BreadcrumbItemProps
}

func (bcp BreadcrumbProps) GetBreadcrumbs() (BreadcrumbItems []breadcrumbitem.BreadcrumbItemProps) {

	var breadcrumbItem *breadcrumbitem.BreadcrumbItemProps
	for iter, breadcrumbItemProps := range bcp.BreadcrumbItems {
		breadcrumbItem = breadcrumbItemProps
		if iter == len(bcp.BreadcrumbItems)-1 {
			breadcrumbItem.IsActive = true
		}
		BreadcrumbItems = append(BreadcrumbItems, *breadcrumbItem)
	}

	return
}

func Show(props BreadcrumbProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, breadcrumbItemProps := range props.GetBreadcrumbs() {
			templ_7745c5c3_Err = breadcrumbitem.Show(breadcrumbItemProps).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ol></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
