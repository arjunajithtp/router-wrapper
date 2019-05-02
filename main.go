package main

import (
	"github.com/arjunajithtp/router-wrapper/router"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	r := router.Config{
		Port: "1234",
		Routes: []router.Route{
			router.Route{
				Method: http.MethodGet,
				Path: "/id",
				Handler: TestHandler1,
			},
			router.Route{
				Method: http.MethodPost,
				Path: "/name",
				Handler: TestHandler2,
			},
		},
	}

	r.NewRouter()
}

func TestHandler1(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello id")
}

func TestHandler2(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello name")
}
