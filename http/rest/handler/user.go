package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/user/model"
	"mitra-kirim-be-mgmt/internal/user/service"
	"mitra-kirim-be-mgmt/pkg/response"
	"net/http"
)

type UserHandler struct {
	Svc *service.User
	Log *logrus.Logger
}

func NewUserHandler(svc *service.User, log *logrus.Logger) *UserHandler {
	return &UserHandler{Svc: svc, Log: log}
}

func (h *UserHandler) Login(c echo.Context) error {
	var req model.LoginRequest

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return response.ErrorBadRequest(c, nil, "Username or password is required")
	}

	req.Username = username
	req.Password = password

	login, err := h.Svc.Login(&req)
	if err != nil {
		return response.ErrorUnauthorized(c, err, "Invalid username or password")
	}
	return response.SuccessOK(c, login)
}

func (h *UserHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Username or password is required")
	}

	login, err := h.Svc.Register(&req)
	if err != nil {
		return response.ErrorUnauthorized(c, err, "Register is failed")
	}
	return response.SuccessOK(c, login)
}

func (h *UserHandler) Refresh(c echo.Context) error {
	var req model.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Bad Request")
	}

	accessToken, err := h.Svc.Refresh(&req)
	if err != nil {
		return response.ErrorUnauthorized(c, err, "Invalid refresh token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": accessToken,
	})
}

func (h *UserHandler) Information(c echo.Context) error {
	userID := c.Get("userID").(string)

	user, err := h.Svc.GetInformation(userID)
	if err != nil {
		return response.ErrorBadRequest(c, err, "Failed")
	}

	return response.SuccessOK(c, user)
}

func (h *UserHandler) Update(c echo.Context) error {
	var req model.UserUpdate

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

	userID := c.Get("userID").(string)

	updatedId, err := h.Svc.Update(&req, userID)
	if err != nil {
		return response.ErrorBadRequest(c, err, "Failed")
	}

	return response.SuccessCreatedReturnId(c, updatedId)
}

func (h *UserHandler) ChangeProfile(c echo.Context) error {
	img, err := c.FormFile("img")
	if err != nil || img.Size == 0 {
		h.Log.Error(err)
		return response.ErrorBadRequest(c, err, "Img not uploaded")
	}

	username := c.Get("username").(string)

	ok, err := h.Svc.UpdatePicture(img, username)
	if err != nil {
		return response.ErrorBadRequest(c, err, "Failed")
	}

	return response.SuccessOK(c, ok)
}
