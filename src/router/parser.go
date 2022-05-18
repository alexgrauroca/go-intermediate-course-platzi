package router

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func parseBody(c echo.Context) (map[string]interface{}, string) {
	var body map[string]interface{}

	tmpBody, ioErr := ioutil.ReadAll(c.Request().Body)

	if ioErr != nil {
		return nil, ioErr.Error()
	}

	if tmpBody == nil {
		return nil, "Body required"
	}

	err := json.Unmarshal(tmpBody, &body)

	if err != nil {
		return nil, err.Error()
	}

	return body, ""
}
