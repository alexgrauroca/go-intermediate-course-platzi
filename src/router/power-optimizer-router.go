package router

import (
	"net/http"
	"power-optimizer/src/optimizer"

	"github.com/labstack/echo/v4"
)

func PowerOptimizerRouter(c echo.Context) error {
	body, err := parseBody(c)

	if body == nil {
		c.String(http.StatusConflict, err)
	}

	response, httpStatus := optimizer.CalculateOptimizedContractedPowers(body)

	return c.JSON(httpStatus, response)
}
