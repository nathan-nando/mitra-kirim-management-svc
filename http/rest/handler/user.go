package handler

import (
	"github.com/labstack/echo/v4"
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
