package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Config is the configration to create new router
type Config struct {
	Port   string
	Routes []Route
}

// Route is the router config
type Route struct {
	Method  string
	Path    string
	Handler func(echo.Context) error
}

// NewRouter creates new instance of router
func (c *Config) NewRouter() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	for _, route := range c.Routes {
		switch route.Method {
		case http.MethodGet:
			e.GET(route.Path, route.Handler)
		case http.MethodPost:
			e.POST(route.Path, route.Handler)
		}
	}
	e.Logger.Fatal(e.Start(":" + c.Port))
}
