package router

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	RoutesV1(e)
}
