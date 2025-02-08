package converter

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetQueryInt(c echo.Context, param string, defaultValue int) int {
	value := c.QueryParam(param)
	if value == "" {
		return defaultValue
	}

	// Convert string to int
	intValue, err := strconv.Atoi(value)
	if err != nil || intValue < 0 {
		return defaultValue
	}

	return intValue
}
