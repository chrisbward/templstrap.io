package main

import (
	"net/http"

	"github.com/chrisbward/templstrap.io/cmd/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Infoln("hello")

	app := echo.New()

	homepageHandler := handlers.HomepageHandler{}
	componentsPageHandler := handlers.ComponentsPageHandler{}

	app.GET("/", homepageHandler.HandlerHomepageShow)
	app.GET("/components", componentsPageHandler.HandlerComponentsPageShow)

	// Use the middleware to serve static files
	app.Use(middleware.Static("/static"))

	// Define a route to handle requests to the static folder
	app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.Dir("./_static")))))

	app.Start(":9980")
}
