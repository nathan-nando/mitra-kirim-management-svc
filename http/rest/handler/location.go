package handler

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/location/model"
	"mitra-kirim-be-mgmt/internal/location/service"
	"mitra-kirim-be-mgmt/pkg/response"
	"net/http"
	"strconv"
)

type LocationHandler struct {
	Svc *service.Location
	Log *logrus.Logger
}

func NewLocationHandler(service *service.Location, logger *logrus.Logger) *LocationHandler {
	return &LocationHandler{Svc: service, Log: logger}
}

func (h *LocationHandler) List(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.Svc.GetAll(ctx, 0, 100)
	if err != nil {
		return response.ErrorInternal(c, err, "Server error")
	}
	return response.SuccessOK(c, res)
}

func (h *LocationHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	
	var req model.LocationRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Bad Request")
	}

	if err := c.Validate(&req); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		var errorResponse []response.ValidationErrorResponse

		for _, fieldError := range validationErrors {
			errorResponse = append(errorResponse, response.ValidationErrorResponse{
				Field:   fieldError.Field(),
				Message: response.GetErrorMessage(fieldError),
			})
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	username := c.Get("username").(string)

	res, err := h.Svc.Create(ctx, &req, username)
	if err != nil {
		fmt.Println("err", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return response.SuccessCreatedReturnId(c, res)
}

func (h *LocationHandler) Update(c echo.Context) error {
	var req model.LocationRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Bad Request")
	}

	if err := c.Validate(&req); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		var errorResponse []response.ValidationErrorResponse

		for _, fieldError := range validationErrors {
			errorResponse = append(errorResponse, response.ValidationErrorResponse{
				Field:   fieldError.Field(),
				Message: response.GetErrorMessage(fieldError),
			})
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	username := c.Get("username").(string)

	res, err := h.Svc.Update(&req, username)
	if err != nil {
		fmt.Println("err", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return response.SuccessOK(c, res)
}

func (h *LocationHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.ErrorBadRequest(c, nil, "Id is required")
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorBadRequest(c, err, "Id is not valid")
	}

	res, err := h.Svc.Delete(idNumber)
	if err != nil {
		fmt.Println("err", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return response.SuccessCreatedReturnId(c, res)
}
