package main

import (
	"github.com/devianwahyu/go_rest/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	controller.UseSubroute(api)

	e.Logger.Fatal(e.Start(":1323"))
}
