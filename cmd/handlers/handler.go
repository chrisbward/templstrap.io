package handlers

import (
	"net/http"

	"github.com/chrisbward/templstrap.io/pkg/components/breadcrumb/breadcrumbitem"
	"github.com/chrisbward/templstrap.io/pkg/components/navbar/navitem"
	"github.com/chrisbward/templstrap.io/testsuite/entities"
	componentspage "github.com/chrisbward/templstrap.io/testsuite/pages/componentspage/view"
	accordioncomponentspage "github.com/chrisbward/templstrap.io/testsuite/pages/examples/accordion/view"
	cardcomponentspage "github.com/chrisbward/templstrap.io/testsuite/pages/examples/card/view"
	formspage "github.com/chrisbward/templstrap.io/testsuite/pages/forms/view"
	homepage "github.com/chrisbward/templstrap.io/testsuite/pages/homepage/view"
	htmxpagemodel "github.com/chrisbward/templstrap.io/testsuite/pages/htmx/model"
	htmxpage "github.com/chrisbward/templstrap.io/testsuite/pages/htmx/view"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	// "github.com/chrisbward/templstrap.io/entities"
)

var homepageLink = breadcrumbitem.BreadcrumbItemProps{
	Name: "home",
	URL:  "/",
}

var siteNavigation = []*navitem.NavItemProps{
	{
		Name:       "Home",
		Link:       "/",
		IsDisabled: true,
	},
	{
		Name:       "Content",
		Link:       "/content",
		IsDisabled: true,
	},
	{
		Name:       "Components",
		Link:       "/components",
		IsDropdown: true,
	},
	{
		Name: "Forms",
		Link: "/forms",
	},
	{
		Name: "htmx",
		Link: "/htmx",
	},
}

type HomepageHandler struct {
}

func (h HomepageHandler) HandlerHomepageShow(c echo.Context) (err error) {

	return render(c, homepage.Show())
}

type ComponentsPageHandler struct {
}

func (h ComponentsPageHandler) HandlerComponentsPageShow(c echo.Context) (err error) {

	pageProps := entities.NewPageProps("components", "Components", siteNavigation, nil, []*breadcrumbitem.BreadcrumbItemProps{
		&homepageLink,
		{
			Name: "components",
			URL:  "/components",
		},
	})

	return render(c, componentspage.Show(pageProps))
}

type FormsPageHandler struct {
}

func (h FormsPageHandler) HandlerFormsPageShow(c echo.Context) (err error) {
	pageProps := entities.NewPageProps("forms", "Forms", siteNavigation, nil, []*breadcrumbitem.BreadcrumbItemProps{
		&homepageLink,
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

	req := c.Request()
	method := req.Method

	pageProps := entities.NewPageProps("htmx", "htmx", siteNavigation, nil, []*breadcrumbitem.BreadcrumbItemProps{
		&homepageLink,
		{
			Name: "htmx",
			URL:  "/htmx",
		},
	})

	htmxpagemodel := htmxpagemodel.NewHTMXPageModel(c, (method == http.MethodPost))
	logrus.Infoln("returned content")
	return render(c, htmxpage.Show(htmxpagemodel, pageProps))
}

type CardComponentsPageHandler struct {
}

func (h CardComponentsPageHandler) HandlerComponentPageShow(c echo.Context) (err error) {

	return render(c, cardcomponentspage.Show())
}

type AccordionComponentsPageHandler struct {
}

func (h AccordionComponentsPageHandler) HandlerComponentPageShow(c echo.Context) (err error) {
	pageProps := entities.NewPageProps(
		"components",
		"Components",
		siteNavigation,
		nil,
		[]*breadcrumbitem.BreadcrumbItemProps{
			&homepageLink,
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
