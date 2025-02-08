package response

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success      bool        `json:"success" example:"true"`
	Message      string      `json:"message" example:"success"`
	ResponseTime int64       `json:"responseTime" example:"100000"`
	RequestID    string      `json:"requestId" example:"requestId"`
	Data         interface{} `json:"data"`
}
type ResponseOnlyId struct {
	Id interface{} `json:"id"`
}
type Data map[string]interface{}

func buildResponseMsg(defaultMsg string, msg ...string) string {
	if len(msg) == 0 {
		return defaultMsg
	}
	var response string
	for i, item := range msg {
		response += item
		if len(msg)-1 != i {
			response += ", "
		}
	}
	return response
}

func Success(c echo.Context, code int, data interface{}, msg ...string) error {

	responseMsg := buildResponseMsg("Success")

	if data == nil {
		data = map[string]interface{}{}
	}

	var latency int64
	if startTime, ok := c.Get("startTime").(time.Time); ok {
		latency = time.Since(startTime).Milliseconds()
	}

	res := Response{
		Success:      true,
		Message:      responseMsg,
		ResponseTime: latency,
		RequestID:    c.Get("requestId").(string),
		Data:         data,
	}
	return c.JSON(code, res)
}
func SuccessOK(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusOK, data, msg...)
}
func SuccessCreated(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusCreated, data, msg...)
}
func SuccessCreatedReturnId(c echo.Context, id interface{}, msg ...string) error {
	onlyId := ResponseOnlyId{Id: id}
	return Success(c, http.StatusCreated, onlyId, msg...)
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "gte":
		return "Value must be greater than or equal to 18"
	case "lte":
		return "Value must be less than or equal to 60"
	default:
		return "Invalid value"
	}
}
