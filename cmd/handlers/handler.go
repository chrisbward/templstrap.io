package handlers

import (
	componentspage "github.com/chrisbward/templstrap.io/testsuite/pages/components/view"
	homepage "github.com/chrisbward/templstrap.io/testsuite/pages/homepage/view"
	"github.com/labstack/echo/v4"
	// "github.com/chrisbward/templstrap.io/entities"
)

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
