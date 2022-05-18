package main

import (
	"fmt"
	"power-optimizer/src/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(
		middleware.Recover(), // Recover from all panics to always have your server up
	)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		fmt.Println("[ERROR]", c.Path(), c.QueryParams(), err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}
	router.Routes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
