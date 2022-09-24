package echo

import (
	"tmp_project_name/app/infrastructures/echo/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func defineRouting(e *echo.Echo) {

	e.GET("/", handlers.Root)
	e.GET("/index", handlers.Index)
	e.POST("/login", handlers.Login)

	apiPath := "/api"
	api := e.Group(apiPath)
	api.Use(middleware.JWT([]byte("secret")))
	api.POST("/refreshToken", handlers.RefreshToken)
}
