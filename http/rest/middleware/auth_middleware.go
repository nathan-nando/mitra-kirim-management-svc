package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"mitra-kirim-be-mgmt/internal/user/model"
	"net/http"
)

func (m *CustomMiddleware) AuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(m.JwtKey),
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims := &model.Claims{}
			token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(m.JwtKey), nil
			})
			m.Log.Infof("TYPE: %v", claims.TokenType)
			if err != nil {
				return nil, err
			}

			if claims.TokenType == "refresh" {
				return nil, jwt.ErrTokenInvalidClaims
			}
			return token, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		},
	})
}
