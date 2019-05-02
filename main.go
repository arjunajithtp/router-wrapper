package main

import (
	"github.com/arjunajithtp/router-wrapper/router"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	r := router.Config{
		Port: "1234",
		Groups: []router.Group{
			router.Group{
				Prefix: "/redfish/v1",
				Middlewares: []echo.MiddlewareFunc{
					middleware.Logger(),
					middleware.Recover(),
				},
				Routes: []router.Route{
					router.Route{
						Method:  http.MethodGet,
						Path:    "/get",
						Handler: GetHandler,
					},
					router.Route{
						Method:  http.MethodPost,
						Path:    "/post/:id",
						Handler: PostHandler,
					},
				},
			},
			router.Group{
				Prefix: "/aggregator",
				Middlewares: []echo.MiddlewareFunc{
					middleware.Logger(),
					middleware.Recover(),
				},
				Routes: []router.Route{
					router.Route{
						Method:  http.MethodGet,
						Path:    "/get",
						Handler: GetHandler,
					},
					router.Route{
						Method:  http.MethodPost,
						Path:    "/post/:id",
						Handler: PostHandler,
					},
				},
			},
		},
	}
	r.NewRouter()
}

// GetHandler is for get method
func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Get")
}

// PostHandler is for post method
func PostHandler(c echo.Context) error {
	data := "Hello Post: " + c.Param("id")
	return c.JSON(http.StatusOK, data)
}
