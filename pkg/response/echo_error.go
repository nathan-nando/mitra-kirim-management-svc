package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	RequestID string `json:"requestId"`
	Success   bool   `json:"success" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode,omitempty"`
	Internal  error  `json:"-"`
}

func errorEcho(c echo.Context, code int, error interface{}, msg ...string) error {
	responseMsg := buildResponseMsg("error", msg...)

	if error == nil {
		error = map[string]interface{}{}
	}

	requestID := "TEST"

	res := Response{
		Success:   true,
		Message:   responseMsg,
		RequestID: requestID,
		Data:      error,
	}
	return c.JSON(code, res)
}

func ErrorInternal(c echo.Context, data interface{}, msg ...string) error {
	return errorEcho(c, http.StatusInternalServerError, data, msg...)
}

func ErrorBadRequest(c echo.Context, data interface{}, msg ...string) error {
	return errorEcho(c, http.StatusBadRequest, data, msg...)
}

func ErrorNotFound(c echo.Context, data interface{}, msg ...string) error {
	return errorEcho(c, http.StatusNotFound, data, msg...)
}

func ErrorForbidden(c echo.Context, data interface{}, msg ...string) error {
	return errorEcho(c, http.StatusForbidden, data, msg...)
}

func ErrorUnauthorized(c echo.Context, data interface{}, msg ...string) error {
	return errorEcho(c, http.StatusUnauthorized, data, msg...)
}
