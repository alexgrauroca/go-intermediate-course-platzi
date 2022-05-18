package router

import "github.com/labstack/echo/v4"

func RoutesV1(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/do-optimize-powers", PowerOptimizerRouter)
}
