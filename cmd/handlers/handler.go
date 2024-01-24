package handlers

import (
	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
	componentspage "github.com/chrisbward/templstrap.io/testsuite/pages/componentspage/view"
	accordioncomponentspage "github.com/chrisbward/templstrap.io/testsuite/pages/examples/accordion/view"
	cardcomponentspage "github.com/chrisbward/templstrap.io/testsuite/pages/examples/card/view"
	formspage "github.com/chrisbward/templstrap.io/testsuite/pages/forms/view"
	homepage "github.com/chrisbward/templstrap.io/testsuite/pages/homepage/view"
	htmxpage "github.com/chrisbward/templstrap.io/testsuite/pages/htmx/view"
	"github.com/labstack/echo/v4"
	// "github.com/chrisbward/templstrap.io/entities"
)

var homepageLink = breadcrumbitem.BreadcrumbItemProps{
	Name: "home",
	URL:  "/",
}

type HomepageHandler struct {
}

func (h HomepageHandler) HandlerHomepageShow(c echo.Context) (err error) {

	return render(c, homepage.Show())
}

type ComponentsPageHandler struct {
}

func (h ComponentsPageHandler) HandlerComponentsPageShow(c echo.Context) (err error) {

	return render(c, componentspage.Show(componentspage.PageProps{
		Title: "Components",
	}))
}

type FormsPageHandler struct {
}

func (h FormsPageHandler) HandlerFormsPageShow(c echo.Context) (err error) {
	pageProps := entities.NewPageProps("forms", nil, nil, []breadcrumbitem.BreadcrumbItemProps{
		homepageLink,
		{
			Name: "forms",
			URL:  "/forms",
		},
	})
	return render(c, formspage.Show(pageProps))
}

type HTMXPageHandler struct {
}

func (h HTMXPageHandler) HandlerHTMXPageShow(c echo.Context) (err error) {
	pageProps := entities.NewPageProps("htmx", nil, nil, []breadcrumbitem.BreadcrumbItemProps{
		homepageLink,
		{
			Name: "htmx",
			URL:  "/htmx",
		},
	})
	return render(c, htmxpage.Show(pageProps))
}

type CardComponentsPageHandler struct {
}

func (h CardComponentsPageHandler) HandlerComponentPageShow(c echo.Context) (err error) {

	return render(c, cardcomponentspage.Show())
}

type AccordionComponentsPageHandler struct {
}

func (h AccordionComponentsPageHandler) HandlerComponentPageShow(c echo.Context) (err error) {
	pageProps := entities.NewPageProps("htmx", nil, nil, []breadcrumbitem.BreadcrumbItemProps{
		homepageLink,
		{
			Name: "components",
			URL:  "/components",
		},
		{
			Name: "accordion",
			URL:  "/components/accordion",
		},
	})
	return render(c, accordioncomponentspage.Show(pageProps))
}
