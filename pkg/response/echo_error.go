package response

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	errorw2 "mitra-kirim-be/mitra-kirim-be-mgmt/pkg/errorw"
	"mitra-kirim-be/mitra-kirim-be-mgmt/pkg/logger"
	"net/http"
	"sort"
)

type ErrorResponse struct {
	HTTPCode  int               `json:"-"`
	Success   bool              `json:"success" example:"false"`
	Message   string            `json:"message"`
	ErrorCode errorw2.ErrorCode `json:"error_code,omitempty"`
	RequestID string            `json:"request_id"`
	Internal  error             `json:"-"`
}

type ErrResponseFunc func(errInner error) ErrorResponse

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by CustomHTTPErrorHandler
func (e ErrorResponse) StatusCode() int {
	return e.HTTPCode
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrInternalServerError creates a new error response representing an internal server error (HTTP 500)
func ErrInternalServerError(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val = errorw2.ErrInternalServer
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusUnauthorized,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

func ErrUnauthorized(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val = errorw2.ErrUnauthorized
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusUnauthorized,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrForbidden creates a new error response representing an authorization failure (HTTP 403)
func ErrForbidden(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val = errorw2.ErrForbidden
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusForbidden,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrSessionExpired creates a new error response representing an session expired error
func ErrSessionExpired(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val = errorw2.ErrSessionExpired
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  440,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrNotFound creates a new error response representing a resource not found (HTTP 404)
func ErrNotFound(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val = errorw2.ErrBadRequest
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusNotFound,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

// ErrBadRequest creates a new error response representing a bad request (HTTP 400)
func ErrBadRequest(err error) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode errorw2.ErrorCode
	var errorMessage string

	var val errorw2.Error
	if errors.As(originalErr, &val) {
		errorCode = val.Code
		errorMessage = val.Message
	}

	return ErrorResponse{
		HTTPCode:  http.StatusBadRequest,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

func HTTPError(err error, statusCode int, errorCode errorw2.ErrorCode, message string) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	return ErrorResponse{
		HTTPCode:  statusCode,
		Message:   message,
		ErrorCode: errorCode,
		Internal:  err,
	}
}

func CustomHttpErrorHandler(log logger.Logger,
	mapErrorResponse map[errorw2.ErrorCode]ErrResponseFunc, withStack bool) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var requestID = logger.GetRequestID(c.Request().Context())

		err = ConvertError(err, mapErrorResponse)

		var errorResponse ErrorResponse
		if !errors.As(err, &errorResponse) {
			errorResponse = ErrorResponse{
				Success:   false,
				HTTPCode:  http.StatusInternalServerError,
				Message:   errorw2.ErrInternalServer.Message,
				ErrorCode: errorw2.ErrInternalServer.Code,
				Internal:  err,
			}
			err = errorResponse
		}
		errorResponse.RequestID = requestID

		// handles resource not found errors
		if errors.Is(errorResponse.Internal, echo.ErrNotFound) {
			err = HTTPError(errorResponse.Internal, http.StatusNotFound, errorw2.ErrResourceNotFound.Code, "requested endpoint is not registered")
		}

		// Handles validation error
		if errors.As(errorResponse.Internal, &validation.Errors{}) || errors.As(errorResponse.Internal, &validation.ErrorObject{}) {
			err = HTTPError(errorResponse.Internal, http.StatusBadRequest, errorw2.ErrBadRequest.Code, errorResponse.Internal.Error())
		}

		if !errors.As(err, &errorResponse) {
			errorResponse = ErrInternalServerError(err)
		}
		errorResponse.RequestID = logger.GetRequestID(c.Request().Context())

		if withStack {
			if sterr, ok := errorResponse.Internal.(stackTracer); ok {
				fmt.Printf("%+v\n", sterr.StackTrace())
			}
		}

		log.Error(errorResponse.Internal)

		errJson := c.JSON(errorResponse.HTTPCode, errorResponse)
		if errJson != nil {
			log.Error(errJson)
		}
	}
}

func ConvertError(err error, mapError map[errorw2.ErrorCode]ErrResponseFunc) error {

	var arrKey []int
	for key, _ := range mapError {
		arrKey = append(arrKey, int(key))
	}
	sort.Slice(arrKey, func(i, j int) bool {
		return arrKey[i] > arrKey[j]
	})

	var val errorw2.Error
	if errors.As(err, &val) {
		for _, key := range arrKey {
			if int(val.Code) >= key {
				return mapError[errorw2.ErrorCode(key)](val)
			}
		}
	}
	return err
}

var MapDefaultErrResponse = map[errorw2.ErrorCode]ErrResponseFunc{
	errorw2.ErrInternalServer.Code:   ErrInternalServerError,
	errorw2.ErrSessionExpired.Code:   ErrSessionExpired,
	errorw2.ErrForbidden.Code:        ErrForbidden,
	errorw2.ErrUnauthorized.Code:     ErrUnauthorized,
	errorw2.ErrBadRequest.Code:       ErrBadRequest,
	errorw2.ErrResourceNotFound.Code: ErrNotFound,
}
