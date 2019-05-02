package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Config is the configration to create new router
type Config struct {
	Port   string
	Groups []Group
}

// Group is for creating groups in router
type Group struct {
	Prefix      string
	Middlewares []echo.MiddlewareFunc
	Routes      []Route
}

// Route is the router config
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// NewRouter creates new instance of router
func (c *Config) NewRouter() {
	e := echo.New()

	for _, group := range c.Groups {
		g := e.Group(group.Prefix)
		// Middleware
		g.Use(group.Middlewares...)

		for _, route := range group.Routes {
			switch route.Method {
			case http.MethodGet:
				g.GET(route.Path, route.Handler)
			case http.MethodPost:
				g.POST(route.Path, route.Handler)
			}
		}
	}

	e.Logger.Fatal(e.Start(":" + c.Port))
}
