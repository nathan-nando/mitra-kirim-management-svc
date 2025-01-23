package handler

import (
	"github.com/golang-jwt/jwt/v5"
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
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}

	if req.Username != "admin" || req.Password != "password" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Generate tokens
	accessToken, err := h.Svc.GenerateAccessToken("123", "admin@example.com")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}

	refreshToken, err := h.Svc.GenerateRefreshToken("123", "admin@example.com")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate refresh token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

}

func (h *UserHandler) Refresh(c echo.Context) error {
	var req model.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}

	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return service.JwtKey, nil
	})

	if !token.Valid {
		h.Log.Println("Token is not valid")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
	}

	if err != nil || !token.Valid || claims.TokenType != "refresh" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid"})
	}

	// Generate a new access token
	accessToken, err := h.Svc.GenerateAccessToken(claims.UserID, claims.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate access token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": accessToken,
	})
}
