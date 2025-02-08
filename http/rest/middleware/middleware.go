package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type CustomMiddleware struct {
	Log    *logrus.Logger
	JwtKey string
}

func NewCustomMiddleware(log *logrus.Logger, jwtKey string) *CustomMiddleware {
	return &CustomMiddleware{Log: log, JwtKey: jwtKey}
}

func (m *CustomMiddleware) DevMode() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := uuid.NewString()
			start := time.Now()

			c.Set("requestId", requestId)
			c.Set("startTime", start)

			err := next(c)
			if err != nil {
				c.Error(err)
			}
			request := c.Request()
			response := c.Response()

			logMessage := fmt.Sprintf(
				"method=%s, path=%s, status=%d, latency=%s, remote_ip=%s, user_agent=%s",
				request.Method,
				request.URL.Path,
				response.Status,
				c.RealIP(),
				request.UserAgent(),
			)

			// Log with Logrus
			m.Log.WithFields(logrus.Fields{
				"method":     request.Method,
				"path":       request.URL.Path,
				"status":     response.Status,
				"remote_ip":  c.RealIP(),
				"user_agent": request.UserAgent(),
			}).Info(logMessage)

			return nil
		}
	}
}
